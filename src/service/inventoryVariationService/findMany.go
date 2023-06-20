package inventoryVariationService

import (
	"adr/backend/src/prisma/db"
	"context"
)

func FindMany(inventoryId string, client *db.PrismaClient, ctx context.Context) ([]db.InventoryVariationModel, error) {

	iV, err := client.InventoryVariation.
		FindMany(
			db.InventoryVariation.InventoryID.Equals(inventoryId),
		).
		Exec(ctx)

	return iV, err

}
