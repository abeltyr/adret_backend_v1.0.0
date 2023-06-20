package sharedRouter

import (
	"adr/backend/src/graphql/model"
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"adr/backend/src/service/productService"
	"adr/backend/src/utils"
	"context"
	"log"
)

func Product(ctx context.Context, id string) (*model.Product, error) {

	client := ctx.Value(srcModel.ConfigKey("client")).(*db.PrismaClient)

	data, err := productService.Find(id, client, ctx)
	if err != nil {
		log.Println("error productService Find", err)
		return nil, err
	}

	product := utils.ProductConverter(data)

	return product, nil
}
