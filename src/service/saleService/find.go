package saleService

import (
	"adr/backend/src/prisma/db"
	"context"
)

func Find(orderID string, InventoryID string, client *db.PrismaClient, ctx context.Context) (*db.SalesModel, error) {

	sales, err := client.Sales.FindUnique(
		db.Sales.OrderIDInventoryID(
			db.Sales.OrderID.Equals(orderID),
			db.Sales.InventoryID.Equals(InventoryID),
		),
	).Exec(ctx)

	return sales, err
}
