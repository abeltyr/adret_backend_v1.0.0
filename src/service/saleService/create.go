package saleService

import (
	"adr/backend/src/graphql/model"
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"context"
)

func Create(input model.SalesInput, inventory *db.InventoryModel, client *db.PrismaClient, ctx context.Context) (*db.SalesModel, error) {

	order := ctx.Value(srcModel.ConfigKey("order")).(*db.OrderModel)

	profit := input.SellingPrice - inventory.InitialPrice

	createdSale, err := client.Sales.CreateOne(
		db.Sales.Amount.Set(input.Amount),
		db.Sales.SellingPrice.Set(input.SellingPrice),
		db.Sales.Profit.Set(profit),
		db.Sales.Inventory.Link(db.Inventory.ID.Equals(input.InventoryID)),
		db.Sales.Order.Link(db.Order.ID.Equals(order.ID)),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return createdSale, nil

}
