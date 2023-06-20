package productRouter

import (
	"adr/backend/src/graphql/model"
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"adr/backend/src/service/productService"
	"adr/backend/src/service/userService"
	"context"
	"encoding/json"
	"log"
)

func FindMany(ctx context.Context, input model.ProductsFilter) ([]*model.Product, error) {

	client := ctx.Value(srcModel.ConfigKey("client")).(*db.PrismaClient)
	clientCtx := ctx.Value(srcModel.ConfigKey("clientCtx")).(context.Context)
	cognitoUser := ctx.Value(srcModel.ConfigKey("user"))

	_, currentCompany, fetchErr := userService.Checking(cognitoUser, "srcModel.Manager", client, clientCtx)
	if fetchErr != nil {
		log.Println("error userService Checking", fetchErr)
		return nil, fetchErr
	}

	ctx = context.WithValue(ctx, srcModel.ConfigKey("currentCompany"), currentCompany)

	productData, err := productService.FindMany(input, client, ctx)
	if err != nil {
		log.Println("error productService FindMany", fetchErr)
		return nil, err
	}

	j, _ := json.Marshal(productData)

	var products []*model.Product
	_ = json.Unmarshal(j, &products)

	return products, nil
}
