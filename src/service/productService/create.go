package productService

import (
	"adr/backend/src/graphql/model"
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"adr/backend/src/service/tableCountService"
	"context"
	"fmt"
	"log"
)

func Create(input model.CreateProductInput, client *db.PrismaClient, ctx context.Context) (*db.ProductModel, error) {

	currentUser := ctx.Value(srcModel.ConfigKey("currentUser")).(*db.UserModel)
	currentCompany := ctx.Value(srcModel.ConfigKey("currentCompany")).(*db.CompanyModel)

	add := 1
	count, err := tableCountService.Create("productCount", &add, client, ctx)
	if err != nil {
		return nil, err
	}

	parameter := []db.ProductSetParam{}

	var branchId string
	var ok bool
	log.Println(currentUser)
	if currentUser != nil {
		parameter = append(parameter,
			db.Product.CreatorID.Set(currentUser.ID),
		)
		branchId, ok = currentUser.BranchID()
		log.Println(branchId, ok)
		if ok {
			parameter = append(parameter,
				db.Product.BranchID.Set(branchId),
			)
		}
	}
	if currentCompany != nil {
		parameter = append(parameter,
			db.Product.CompanyID.Set(currentCompany.ID),
		)
	}

	productCode := fmt.Sprintf(`%v-%d`, currentCompany.CompanyCode, (1000 + int64(count.Count)))

	createdProduct, err := client.Product.CreateOne(
		db.Product.ProductCode.Set(productCode),
		db.Product.Title.Set(input.Title),
		db.Product.Detail.Set(input.Detail),

		parameter[:]...,
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return createdProduct, nil

}
