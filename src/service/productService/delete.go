package productService

import (
	"adr/backend/src/prisma/db"
	"context"
)

func Delete(id string, client *db.PrismaClient, ctx context.Context) (*db.ProductModel, error) {

	product, err := client.Product.
		FindUnique(
			db.Product.ID.Equals(id),
		).Delete().
		Exec(ctx)

	return product, err
}
