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
	"adr/backend/src/service/userService"
	"adr/backend/src/utils"
	"adr/backend/src/utils/s3Config"
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func Create(ctx context.Context, input model.CreateProductInput) (*model.Product, error) {

	if len(input.Media) > 5 {
		return nil, errors.New("max allowed media is 5")
	}

	client := ctx.Value(srcModel.ConfigKey("client")).(*db.PrismaClient)
	clientCtx := ctx.Value(srcModel.ConfigKey("clientCtx")).(context.Context)
	cognitoUser := ctx.Value(srcModel.ConfigKey("user"))

	currentUser, currentCompany, fetchErr := userService.Checking(cognitoUser, srcModel.Manager, client, clientCtx)
	if fetchErr != nil {
		log.Println("error userService Checking", fetchErr)
		return nil, fetchErr
	}

	ctx = context.WithValue(ctx, srcModel.ConfigKey("currentUser"), currentUser)
	ctx = context.WithValue(ctx, srcModel.ConfigKey("currentCompany"), currentCompany)

	productData, err := productService.Create(input, client, ctx)
	if err != nil {
		log.Println("error productService create", err)
		return nil, err
	}

	ctx = context.WithValue(ctx, srcModel.ConfigKey("product"), productData)

	files := []*db.FileModel{}
	keys := []*s3.ObjectIdentifier{}

	hasMediaError := false
	mediaErrorMessage := "error uploading media"

	for index, media := range input.Media {
		if media != nil {
			key, err := utils.ImageProcessor(*media, "product", productData.ID, index)

			if err != nil {
				hasMediaError = true
				mediaErrorMessage = "error uploading media on" + strconv.Itoa(index)
				break
			}

			keys = append(keys, &s3.ObjectIdentifier{Key: aws.String(key)})

			size := fmt.Sprint(media.Size)

			file, err := fileService.Create(srcModel.UploadedData{
				Name:       media.Filename,
				Uploader:   currentUser.ID,
				URL:        key,
				PreviewURL: key,
				Size:       size,
			}, client, ctx)

			if err != nil {
				hasMediaError = true
				mediaErrorMessage = "fileService creation failed"
				break
			}
			files = append(files, file)

			_, err = fileRelationService.Create(srcModel.UploadedDataRelation{
				File:  file.ID,
				Table: productData.ID,
				Value: "Media",
				Order: index,
			}, client, ctx)
			if err != nil {
				hasMediaError = true
				mediaErrorMessage = "file relation creation failed"
				break
			}
		}

	}
	if hasMediaError {
		productService.Delete(productData.ID, client, ctx)
		if len(keys) > 0 {
			s3Config.MultipleDelete(keys)
		}
		return nil, errors.New(mediaErrorMessage)
	}

	hasInventoryCategoryError := false

	inventoryCategoryError := "inventory creation failed"
	for _, inventory := range input.Inventory {

		inventoryData, err := inventoryService.Create(*inventory, client, ctx)
		if err != nil {
			hasInventoryCategoryError = true
			break
		}

		_, err = priceHistoryService.Create(
			model.CreateInventoryInput{
				InitialPrice:              inventoryData.InitialPrice,
				MinSellingPriceEstimation: inventoryData.MinSellingPriceEstimation,
				MaxSellingPriceEstimation: inventoryData.MaxSellingPriceEstimation,
			}, inventoryData.ID, client, ctx,
		)

		if err != nil {
			hasInventoryCategoryError = true
			inventoryCategoryError = "error on priceHistoryService creation"
			break
		}

		count := 0
		for _, variation := range inventory.Variation {
			productVariation, err := productVariationService.Create(*variation.Title, count, *productData, client, ctx)
			if err != nil {
				hasInventoryCategoryError = true
				inventoryCategoryError = "error on productVariation creation"
				break
			}
			_, err = inventoryVariationService.Create(*variation.Data, *inventoryData, *productVariation, client, ctx)
			if err != nil {
				hasInventoryCategoryError = true
				inventoryCategoryError = "error on inventory variation Service creation"
				break
			}
			count++

		}

		if inventory.Media != nil && len(files) > *inventory.Media {
			_, err = fileRelationService.Create(srcModel.UploadedDataRelation{
				File:  files[*inventory.Media].ID,
				Table: inventoryData.ID,
				Value: "Media",
				Order: 0,
			}, client, ctx)
			if err != nil {
				hasInventoryCategoryError = true
				inventoryCategoryError = "error on inventory media setup"
				break
			}
		}

	}
	if !hasInventoryCategoryError {
		_, err = categoryService.Create(input.Category, client, ctx)
		if err != nil {
			hasInventoryCategoryError = true
			inventoryCategoryError = "error on category service creation"
		}
		_, err = categoryProductService.Create(input.Category, client, ctx)
		if err != nil {
			hasInventoryCategoryError = true
			inventoryCategoryError = "error on category product service creation"
		}
		_, err = categoryCompanyService.Create(input.Category, client, ctx)
		if err != nil {
			hasInventoryCategoryError = true
			inventoryCategoryError = "error on category company service creation"
		}
	}

	if hasInventoryCategoryError {
		productService.Delete(productData.ID, client, ctx)
		if len(keys) > 0 {
			s3Config.MultipleDelete(keys)
		}
		return nil, errors.New(inventoryCategoryError)
	}

	product := utils.ProductConverter(productData)
	return product, nil

}
