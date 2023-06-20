package inventoryService

import (
	"adr/backend/src/graphql/model"
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"context"
)

func Create(input model.CreateInventoryInput, client *db.PrismaClient, ctx context.Context) (*db.InventoryModel, error) {

	product := ctx.Value(srcModel.ConfigKey("product")).(*db.ProductModel)

	createdInventory, err := client.Inventory.CreateOne(
		db.Inventory.InitialPrice.Set(float64(input.InitialPrice)),
		db.Inventory.MinSellingPriceEstimation.Set(float64(input.MinSellingPriceEstimation)),
		db.Inventory.MaxSellingPriceEstimation.Set(float64(input.MaxSellingPriceEstimation)),
		db.Inventory.Product.Link(db.Product.ID.Equals(product.ID)),
		db.Inventory.Available.Set(input.Amount),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return createdInventory, nil

}
