package categoryProductService

import (
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"context"
	"strings"
)

func Update(categoryName string, client *db.PrismaClient, ctx context.Context) (bool, error) {

	product := ctx.Value(srcModel.ConfigKey("product")).(*db.ProductModel)

	category := strings.ToLower(categoryName)
	_, err := client.CategoryProduct.FindMany(
		db.CategoryProduct.ProductID.Equals(product.ID),
	).Update(
		db.CategoryProduct.CategoryName.Set(category),
	).Exec(ctx)

	if err != nil {
		return false, err
	}

	return true, err
}
