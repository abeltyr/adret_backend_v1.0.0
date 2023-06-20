package orderService

import (
	"adr/backend/src/graphql/model"
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"adr/backend/src/utils"
	"context"
	"time"
)

func FindMany(filter model.OrdersFilter, client *db.PrismaClient, ctx context.Context) ([]db.OrderModel, error) {

	var limit int32 = 20

	currentCompany := ctx.Value(srcModel.ConfigKey("currentCompany")).(*db.CompanyModel)

	parameter := []db.OrderWhereParam{
		db.Order.CompanyID.Equals(currentCompany.ID),
	}

	if filter.Filter != nil && filter.Filter.Limit != nil {
		limit = utils.LimitSetter(int32(*filter.Filter.Limit))
	}

	if filter.SellerID != nil {
		parameter = append(parameter,
			db.Order.SellerID.Equals(*filter.SellerID),
		)
	}

	if filter.MinTotalPrice != nil {
		parameter = append(parameter,
			db.Order.TotalPrice.Gte(*filter.MinTotalPrice),
		)
	}

	if filter.MaxTotalPrice != nil {
		parameter = append(parameter,
			db.Order.TotalPrice.Lte(*filter.MaxTotalPrice),
		)
	}
	if filter.StartDate != nil {
		data, _ := time.Parse("2006-01-02", *filter.StartDate)
		parameter = append(parameter,
			db.Order.Date.Gte(data),
		)
	}

	if filter.EndDate != nil {
		data, _ := time.Parse("2006-01-02", *filter.EndDate)

		parameter = append(parameter,
			db.Order.Date.Lte(data),
		)
	}

	FetchOrder := client.Order.
		FindMany(
			parameter[:]...,
		).
		Take(int(limit)).OrderBy(
		db.Order.CreatedAt.Order(db.DESC),
	)

	if filter.Filter != nil {
		if filter.Filter.After != nil {
			FetchOrder = client.Order.
				FindMany(
					parameter[:]...,
				).
				Take(int(limit)).
				Skip(1).
				Cursor(db.Order.ID.Cursor(string(*filter.Filter.After))).
				OrderBy(
					db.Order.CreatedAt.Order(db.DESC),
				)
		}

		if filter.Filter.Before != nil {
			FetchOrder = client.Order.
				FindMany(
					parameter[:]...,
				).
				Take(int(limit)).
				Skip(1).
				Cursor(db.Order.ID.Cursor(string(*filter.Filter.Before))).
				OrderBy(
					db.Order.CreatedAt.Order(db.ASC),
				)
		}
	}
	sales, err := FetchOrder.
		Exec(ctx)

	return sales, err
}
