package inventoryVariationService

import (
	"adr/backend/src/prisma/db"
	"context"
	"strings"
)

func Create(data string, inventory db.InventoryModel, productVariation db.ProductVariationModel, client *db.PrismaClient, ctx context.Context) (*db.InventoryVariationModel, error) {
	title := strings.ToLower(data)
	createdCategoryCompany, err := client.InventoryVariation.
		UpsertOne(
			db.InventoryVariation.InventoryIDProductVariationID(
				db.InventoryVariation.InventoryID.Equals(inventory.ID),
				db.InventoryVariation.ProductVariationID.Equals(productVariation.ID),
			),
		).
		Create(
			db.InventoryVariation.Data.Set(title),
			db.InventoryVariation.Inventory.Link(db.Inventory.ID.Equals(inventory.ID)),
			db.InventoryVariation.ProductVariation.Link(db.ProductVariation.ID.Equals(productVariation.ID)),
		).
		Update(
			db.InventoryVariation.Data.Set(title),
			db.InventoryVariation.Inventory.Link(db.Inventory.ID.Equals(inventory.ID)),
			db.InventoryVariation.ProductVariation.Link(db.ProductVariation.ID.Equals(productVariation.ID)),
		).
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	return createdCategoryCompany, nil

}
