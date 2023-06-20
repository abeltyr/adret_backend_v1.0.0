package productService

import (
	"adr/backend/src/prisma/db"
	"context"
)

func FindByCode(productCode string, client *db.PrismaClient, ctx context.Context) (*db.ProductModel, error) {

	product, err := client.Product.
		FindUnique(
			db.Product.ProductCode.Equals(productCode),
		).
		Exec(ctx)

	return product, err
}
