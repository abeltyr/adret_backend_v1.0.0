package productService

import (
	"adr/backend/src/graphql/model"
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"context"
)

func Update(input model.UpdateProductInput, client *db.PrismaClient, ctx context.Context) (*db.ProductModel, error) {

	product := ctx.Value(srcModel.ConfigKey("product")).(*db.ProductModel)

	title := product.Title
	detail := product.Detail

	if input.Title != nil {
		title = *input.Title
	}

	if input.Detail != nil {
		detail = *input.Detail
	}

	updated, err := client.Product.FindUnique(
		db.Product.ID.Equals(input.ID),
	).Update(
		db.Product.Title.Set(title),
		db.Product.Detail.Set(detail),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return updated, err
}
