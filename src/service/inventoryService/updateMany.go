package inventoryService

import (
	"adr/backend/src/graphql/model"
	"adr/backend/src/prisma/db"
	"context"
)

func UpdateMany(input model.UpdateInventoryInput, client *db.PrismaClient, ctx context.Context) (bool, error) {

	parameter := []db.InventorySetParam{}

	if input.InitialPrice != nil {
		parameter = append(parameter, db.Inventory.InitialPrice.Set(*input.InitialPrice))

	}

	if input.MinSellingPriceEstimation != nil {
		parameter = append(parameter, db.Inventory.MinSellingPriceEstimation.Set(*input.MinSellingPriceEstimation))
	}

	if input.MaxSellingPriceEstimation != nil {
		parameter = append(parameter, db.Inventory.MaxSellingPriceEstimation.Set(*input.MaxSellingPriceEstimation))
	}

	_, err := client.Inventory.FindMany(
		db.Inventory.ProductID.Equals(input.ProductID),
	).Update(
		parameter[:]...,
	).Exec(ctx)

	if err != nil {
		return false, err
	}

	return true, err
}
