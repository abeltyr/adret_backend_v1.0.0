package productService

import (
	"adr/backend/src/prisma/db"
	"context"
)

func Find(id string, client *db.PrismaClient, ctx context.Context) (*db.ProductModel, error) {

	product, err := client.Product.
		FindUnique(
			db.Product.ID.Equals(id),
		).
		Exec(ctx)

	return product, err
}
