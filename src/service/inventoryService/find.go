package inventoryService

import (
	"adr/backend/src/prisma/db"
	"context"
)

func Find(id string, client *db.PrismaClient, ctx context.Context) (*db.InventoryModel, error) {

	inventory, err := client.Inventory.FindUnique(
		db.Inventory.ID.Equals(id),
	).Exec(ctx)

	return inventory, err
}
