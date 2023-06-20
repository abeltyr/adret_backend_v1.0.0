package categoryProductService

import (
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"context"
)

func Delete(client *db.PrismaClient, ctx context.Context) (bool, error) {

	product := ctx.Value(srcModel.ConfigKey("product")).(*db.ProductModel)

	_, err := client.CategoryProduct.FindMany(
		db.CategoryProduct.ProductID.Equals(product.ID),
	).Delete().Exec(ctx)

	if err != nil {
		return false, err
	}

	return true, err
}
