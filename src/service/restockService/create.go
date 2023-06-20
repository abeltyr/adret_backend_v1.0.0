package restockService

import (
	"adr/backend/src/prisma/db"
	"context"
)

func Create(amount int, inventoryId string, client *db.PrismaClient, ctx context.Context) (*db.RestockModel, error) {
	restock, err := client.Restock.CreateOne(
		db.Restock.Amount.Set(amount),
		db.Restock.Inventory.Link(db.Inventory.ID.Equals((inventoryId))),
	).Exec(ctx)

	return restock, err

}
