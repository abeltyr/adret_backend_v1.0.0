package productRouter

import (
	"adr/backend/src/graphql/model"
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"adr/backend/src/service/categoryCompanyService"
	"adr/backend/src/service/categoryProductService"
	"adr/backend/src/service/categoryService"
	"adr/backend/src/service/fileRelationService"
	"adr/backend/src/service/fileService"
	"adr/backend/src/service/inventoryService"
	"adr/backend/src/service/inventoryVariationService"
	"adr/backend/src/service/priceHistoryService"
	"adr/backend/src/service/productService"
	"adr/backend/src/service/productVariationService"
	"adr/backend/src/service/restockService"
	"adr/backend/src/service/userService"
	"adr/backend/src/utils"
	"context"
	"errors"
	"fmt"
	"log"
)

func Update(ctx context.Context, input model.UpdateProductInput) (*model.Product, error) {

	if len(input.Media) > 5 {
		return nil, errors.New("max allowed media is 5")
	}

	client := ctx.Value(srcModel.ConfigKey("client")).(*db.PrismaClient)
	clientCtx := ctx.Value(srcModel.ConfigKey("clientCtx")).(context.Context)
	cognitoUser := ctx.Value(srcModel.ConfigKey("user"))

	currentUser, company, fetchErr := userService.Checking(cognitoUser, srcModel.Manager, client, clientCtx)
	if fetchErr != nil {
		log.Println("error userService Checking", fetchErr)
		return nil, fetchErr
	}

	fetchProduct, err := productService.Find(input.ID, client, clientCtx)

	ctx = context.WithValue(ctx, srcModel.ConfigKey("product"), fetchProduct)
	ctx = context.WithValue(ctx, srcModel.ConfigKey("currentCompany"), company)
	if err != nil {
		log.Println("error productService Find", err)
		return nil, err
	}
	companyId, ok := fetchProduct.CompanyID()

	if !ok || companyId != company.ID {
		return nil, errors.New("trying to update another company product")
	}

	productData, err := productService.Find(input.ID, client, ctx)
	if err != nil {
		log.Println("error productService Find", err)
		return nil, err
	}
	if (input.Title != nil && productData.Title != *input.Title) || (input.Detail != nil && productData.Detail != *input.Detail) {
		productData, err = productService.Update(input, client, ctx)
		if err != nil {
			log.Println("error productService Update", err)
			return nil, err
		}
	}

	ctx = context.WithValue(ctx, srcModel.ConfigKey("product"), productData)

	fileRelations, err := fileRelationService.FindMany(productData.ID, client, ctx)
	if err != nil {
		log.Println("fileRelationService.FindMany", err)
		return nil, err
	}

	if len(fileRelations) > len(input.Media) {
		for i := len(input.Media); i < len(fileRelations); i++ {
			_, err = fileRelationService.Delete(srcModel.UploadedDataRelation{
				Table: productData.ID,
				Value: "Media",
				Order: i,
			}, client, ctx)
			if err != nil {
				log.Println("fileRelationService.Create", err)
				return nil, err
			}
		}
	}

	files := []*db.FileModel{}
	for index, media := range input.Media {
		var fileData = media.File
		var key string
		size := ""
		if fileData != nil {
			key, err = utils.ImageProcessor(*media.File, "product", productData.ID, index)
			if err != nil {
				return nil, err
			}
			size = fmt.Sprint(fileData.Size)
		} else {
			key = *media.URL
		}

		fetchedFiles, err := fileService.FindByUrl(
			key, client, ctx)

		var file *db.FileModel

		if len(fetchedFiles) == 0 || err != nil {
			file, err = fileService.Create(srcModel.UploadedData{
				Name:       fileData.Filename,
				Uploader:   currentUser.ID,
				URL:        key,
				PreviewURL: key,
				Size:       size,
			}, client, ctx)
			if err != nil {
				return nil, err
			}
		} else if len(fetchedFiles) > 0 {
			file = &fetchedFiles[0]
		}

		files = append(files, file)

		_, err = fileRelationService.Create(srcModel.UploadedDataRelation{
			File:  file.ID,
			Table: productData.ID,
			Value: "Media",
			Order: index,
		}, client, ctx)
		if err != nil {
			log.Println("fileRelationService.Create", err)
			return nil, err
		}

	}
	for _, inventory := range input.Inventory {
		var inventoryData *db.InventoryModel
		if inventory.ID == nil {

			inventoryData, err = inventoryService.Create(
				model.CreateInventoryInput{
					Amount:                    *inventory.Amount,
					InitialPrice:              *inventory.InitialPrice,
					MinSellingPriceEstimation: *inventory.MinSellingPriceEstimation,
					MaxSellingPriceEstimation: *inventory.MaxSellingPriceEstimation,
				}, client, ctx,
			)

			if err != nil {
				log.Println("error inventoryService create", err)
				return nil, err
			}

		} else {
			fetchInventory, err := inventoryService.Find(*inventory.ID, client, ctx)
			if inventory.Amount != nil && fetchInventory.SalesAmount > *inventory.Amount {
				log.Println("error inventoryService create", err)
				return nil, errors.New("can't updated sold inventory")
			}

			inventoryData, err = inventoryService.Update(
				model.UpdateProductInventoryInput{
					Amount:                    inventory.Amount,
					InitialPrice:              inventory.InitialPrice,
					MinSellingPriceEstimation: inventory.MinSellingPriceEstimation,
					MaxSellingPriceEstimation: inventory.MaxSellingPriceEstimation,
					ID:                        inventory.ID,
				}, client, ctx)

			if err != nil {
				log.Println("error inventoryService create", err)
				return nil, err
			}

			// if their is a price update we save that to price history
			if inventoryData.InitialPrice != fetchInventory.InitialPrice || inventoryData.MaxSellingPriceEstimation != fetchInventory.MaxSellingPriceEstimation || inventoryData.MinSellingPriceEstimation != fetchInventory.MinSellingPriceEstimation {
				_, err = priceHistoryService.Create(
					model.CreateInventoryInput{
						InitialPrice:              inventoryData.InitialPrice,
						MinSellingPriceEstimation: inventoryData.MinSellingPriceEstimation,
						MaxSellingPriceEstimation: inventoryData.MaxSellingPriceEstimation,
					}, inventoryData.ID, client, ctx,
				)

				if err != nil {
					log.Println("error priceHistoryService create", err)
					return nil, err
				}

			}

			// if the available is increase we save the restocking
			if inventoryData.Available > fetchInventory.Available {
				_, err = restockService.Create(inventoryData.Available-fetchInventory.Available, inventoryData.ID, client, ctx)
				if err != nil {
					log.Println("error restockService create", err)
					return nil, err
				}
			}

		}

		if inventory.Media != nil && len(files) > *inventory.Media {
			_, err = fileRelationService.Create(srcModel.UploadedDataRelation{
				File:  files[*inventory.Media].ID,
				Table: inventoryData.ID,
				Value: "Media",
				Order: 0,
			}, client, ctx)
			if err != nil {
				log.Println("error fileRelationService create", err)
				return nil, err
			}
		} else if inventory.Media == nil {
			fileRelationFetched, _ := fileRelationService.Find(srcModel.UploadedDataRelation{
				Table: inventoryData.ID,
				Value: "Media",
				Order: 0,
			}, client, ctx)
			if fileRelationFetched != nil {
				_, err = fileRelationService.Delete(srcModel.UploadedDataRelation{
					Table: inventoryData.ID,
					Value: "Media",
					Order: 0,
				}, client, ctx)
				if err != nil {
					log.Println("error fileRelationService delete", err)
					return nil, err
				}
			}
		}

		count := 0
		for _, variation := range inventory.Variation {
			productVariation, err := productVariationService.Create(*variation.Title, count, *productData, client, ctx)
			if err != nil {
				log.Println("error productVariation create", err)
				return nil, err
			}
			_, err = inventoryVariationService.Create(*variation.Data, *inventoryData, *productVariation, client, ctx)
			if err != nil {
				log.Println("error inventoryService create", err)
				return nil, err
			}
			count++
		}

	}

	if input.Category != nil {
		_, err = categoryService.Create(*input.Category, client, ctx)
		if err != nil {
			log.Println("error categoryService create", err)
			return nil, err
		}
		_, err = categoryProductService.Delete(client, ctx)
		if err != nil {
			log.Println("error categoryProductService delete", err)
			return nil, err
		}
		_, err = categoryProductService.Create(*input.Category, client, ctx)
		if err != nil {
			log.Println("error categoryProductService ", err)
			return nil, err
		}
		_, err = categoryCompanyService.Create(*input.Category, client, ctx)
		if err != nil {
			log.Println("error categoryCompanyService create", err)
			return nil, err
		}
	}

	product := utils.ProductConverter(productData)
	return product, nil
}
