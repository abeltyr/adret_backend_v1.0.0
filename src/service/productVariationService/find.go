package productVariationService

import (
	"adr/backend/src/prisma/db"
	"context"
)

func Find(id string, client *db.PrismaClient, ctx context.Context) (*db.ProductVariationModel, error) {

	productVariation, _ := client.ProductVariation.
		FindUnique(
			db.ProductVariation.ID.Equals(id),
		).Exec(ctx)

	return productVariation, nil

}
