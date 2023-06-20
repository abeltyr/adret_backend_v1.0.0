package priceHistoryService

import (
	"adr/backend/src/graphql/model"
	"adr/backend/src/prisma/db"
	"context"
)

func Create(input model.CreateInventoryInput, inventoryId string, client *db.PrismaClient, ctx context.Context) (*db.PriceHistoryModel, error) {

	priceHistory, err := client.PriceHistory.CreateOne(
		db.PriceHistory.Inventory.Link(db.Inventory.ID.Equals((inventoryId))),
		db.PriceHistory.InitialPrice.Set(float64(input.InitialPrice)),
		db.PriceHistory.MinSellingPriceEstimation.Set(float64(input.MinSellingPriceEstimation)),
		db.PriceHistory.MaxSellingPriceEstimation.Set(float64(input.MaxSellingPriceEstimation)),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return priceHistory, nil

}
