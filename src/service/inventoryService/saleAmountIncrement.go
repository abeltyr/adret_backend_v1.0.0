package inventoryService

import (
	"adr/backend/src/prisma/db"
	"context"
)

func SaleAmountIncrement(id string, amount int, client *db.PrismaClient, ctx context.Context) (bool, error) {

	_, err := client.Inventory.FindUnique(
		db.Inventory.ID.Equals(id),
	).Update(
		db.Inventory.SalesAmount.Increment(amount),
	).Exec(ctx)

	if err != nil {
		return false, err
	}

	return true, err
}
