package productRouter

import (
	"adr/backend/src/graphql/model"
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"adr/backend/src/service/productService"
	"adr/backend/src/service/userService"
	"adr/backend/src/utils"
	"context"
	"log"
)

func FindByCode(ctx context.Context, productCode string) (*model.Product, error) {

	client := ctx.Value(srcModel.ConfigKey("client")).(*db.PrismaClient)
	clientCtx := ctx.Value(srcModel.ConfigKey("clientCtx")).(context.Context)
	cognitoUser := ctx.Value(srcModel.ConfigKey("user"))

	_, _, fetchErr := userService.Checking(cognitoUser, "", client, clientCtx)
	if fetchErr != nil {
		log.Println("error userService Checking", fetchErr)
		return nil, fetchErr
	}

	productData, err := productService.FindByCode(productCode, client, ctx)
	if err != nil {
		log.Println("error productService FindByCode", err)
		return nil, err
	}

	product := utils.ProductConverter(productData)
	return product, nil
}
