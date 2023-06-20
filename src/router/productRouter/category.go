package productRouter

import (
	"adr/backend/src/graphql/model"
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"adr/backend/src/service/categoryProductService"
	"context"
	"log"
)

func Category(ctx context.Context, product *model.Product) (*string, error) {

	client := ctx.Value(srcModel.ConfigKey("client")).(*db.PrismaClient)

	productData, err := categoryProductService.Find(srcModel.CategoryProductInput{
		ProductId: product.ID,
	}, client, ctx)
	if err != nil {
		log.Println("error category productService Find", err)
		return nil, err
	}
	if productData == nil {
		return nil, nil
	}
	return &productData.CategoryName, nil
}
