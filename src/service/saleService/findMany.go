package saleService

import (
	"adr/backend/src/graphql/model"
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"adr/backend/src/utils"
	"context"
)

func FindMany(filter model.SalesFilter, client *db.PrismaClient, ctx context.Context) ([]db.SalesModel, error) {

	var limit int32 = 20

	currentCompany := ctx.Value(srcModel.ConfigKey("currentCompany"))

	parameter := []db.SalesWhereParam{}

	if filter.Filter != nil && filter.Filter.Limit != nil {
		limit = utils.LimitSetter(int32(*filter.Filter.Limit))
	}

	if filter.InventoryID != nil {
		parameter = append(parameter,
			db.Sales.InventoryID.Equals(*filter.InventoryID),
		)
	}
	if currentCompany != nil {
		parameter = append(parameter,
			db.Sales.Order.Where(
				db.Order.CompanyID.Equals(currentCompany.(*db.CompanyModel).ID),
			),
		)
	}
	if filter.OrderID != nil {
		parameter = append(parameter,
			db.Sales.OrderID.Equals(*filter.OrderID),
		)
	}

	if filter.MinSellingPrice != nil {
		parameter = append(parameter,
			db.Sales.SellingPrice.Gte(*filter.MinSellingPrice),
		)
	}

	if filter.MaxSellingPrice != nil {
		parameter = append(parameter,
			db.Sales.SellingPrice.Lte(*filter.MaxSellingPrice),
		)
	}

	FetchSales := client.Sales.
		FindMany(
			parameter[:]...,
		).
		Take(int(limit)).OrderBy()

	if filter.Filter != nil {
		if filter.Filter.After != nil {
			FetchSales = client.Sales.
				FindMany(
					parameter[:]...,
				).
				Take(int(limit)).
				Skip(1).
				Cursor(db.Sales.InventoryID.Cursor(string(*filter.Filter.After))).
				OrderBy(
					db.Sales.InventoryID.Order(db.DESC),
				)
		}

		if filter.Filter.Before != nil {
			FetchSales = client.Sales.
				FindMany(
					parameter[:]...,
				).
				Take(int(limit)).
				Skip(1).
				Cursor(db.Sales.InventoryID.Cursor(string(*filter.Filter.Before))).
				OrderBy(
					db.Sales.InventoryID.Order(db.ASC),
				)
		}
	}
	sales, err := FetchSales.
		Exec(ctx)

	return sales, err
}
