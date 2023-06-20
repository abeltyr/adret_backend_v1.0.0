package inventoryService

import (
	"adr/backend/src/graphql/model"
	"adr/backend/src/prisma/db"
	"context"
	"errors"
)

func Update(input model.UpdateProductInventoryInput, client *db.PrismaClient, ctx context.Context) (*db.InventoryModel, error) {

	if input.ID == nil {
		return nil, errors.New("provide an inventory id")
	}

	parameter := []db.InventorySetParam{}

	if input.Amount != nil {
		parameter = append(parameter, db.Inventory.Available.Set(*input.Amount))
	}

	if input.InitialPrice != nil {
		parameter = append(parameter, db.Inventory.InitialPrice.Set(*input.InitialPrice))
	}

	if input.MinSellingPriceEstimation != nil {
		parameter = append(parameter, db.Inventory.MinSellingPriceEstimation.Set(*input.MinSellingPriceEstimation))
	}

	if input.MaxSellingPriceEstimation != nil {
		parameter = append(parameter, db.Inventory.MaxSellingPriceEstimation.Set(*input.MaxSellingPriceEstimation))
	}

	data, err := client.Inventory.FindUnique(
		db.Inventory.ID.Equals(*input.ID),
	).Update(
		parameter[:]...,
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return data, err
}
