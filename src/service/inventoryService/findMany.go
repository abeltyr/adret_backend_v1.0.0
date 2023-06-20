package inventoryService

import (
	"adr/backend/src/graphql/model"
	"adr/backend/src/prisma/db"
	"adr/backend/src/utils"
	"context"
	"time"
)

func FindMany(filter model.InventoriesFilter, client *db.PrismaClient, ctx context.Context) ([]db.InventoryModel, error) {

	var limit int32 = 20

	parameter := []db.InventoryWhereParam{}

	if filter.Filter != nil && filter.Filter.Limit != nil {
		limit = utils.LimitSetter(int32(*filter.Filter.Limit))
	}

	if filter.ProductID != nil {
		parameter = append(parameter,
			db.Inventory.ProductID.Equals(*filter.ProductID),
		)
	}

	if filter.BoughtDates != nil {
		start, _ := time.Parse("2006-01-02", *filter.BoughtDates)
		end := time.Unix(0, (start.UnixMilli()+86400000)*int64(time.Millisecond))

		parameter = append(parameter,
			db.Inventory.CreatedAt.Gte(start),
			db.Inventory.CreatedAt.Lte(end),
		)
	}

	FetchInventory := client.Inventory.
		FindMany(
			parameter[:]...,
		).
		Take(int(limit)).OrderBy(
		db.Inventory.CreatedAt.Order(db.DESC),
	)

	if filter.Filter != nil {
		if filter.Filter.After != nil {
			FetchInventory = client.Inventory.
				FindMany(
					parameter[:]...,
				).
				Take(int(limit)).
				Skip(1).
				Cursor(db.Inventory.ID.Cursor(string(*filter.Filter.After))).
				OrderBy(
					db.Inventory.CreatedAt.Order(db.DESC),
				)
		}

		if filter.Filter.Before != nil {
			FetchInventory = client.Inventory.
				FindMany(
					parameter[:]...,
				).
				Take(int(limit)).
				Skip(1).
				Cursor(db.Inventory.ID.Cursor(string(*filter.Filter.Before))).
				OrderBy(
					db.Inventory.CreatedAt.Order(db.ASC),
				)
		}
	}
	users, err := FetchInventory.
		Exec(ctx)

	return users, err

}
