package orderService

import (
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"context"
	"errors"
	"math"
	"time"
)

func Summary(startDate time.Time, endDate time.Time, sellerId *string, client *db.PrismaClient, ctx context.Context) (float64, float64, error) {

	currentCompany := ctx.Value(srcModel.ConfigKey("currentCompany")).(*db.CompanyModel)

	parameter := []db.OrderWhereParam{
		db.Order.CompanyID.Equals(currentCompany.ID),
		db.Order.Date.Gte(startDate),
		db.Order.Date.Lte(endDate),
	}

	if sellerId != nil {
		parameter = append(parameter,
			db.Order.SellerID.Equals(*sellerId),
		)
	}

	orders, err := client.Order.
		FindMany(
			parameter[:]...,
		).
		Exec(ctx)

	var earning float64 = 0
	var profit float64 = 0

	if len(orders) == 0 {
		return earning, profit, nil
	}

	if err != nil {
		return earning, profit, err
	}

	noValue := true
	for _, order := range orders {
		noValue = false
		earning = earning + order.TotalPrice
		profit = profit + order.TotalProfit

	}

	if noValue {
		return 0, 0, errors.New("no inventory found set")
	}

	return math.Round(earning*100) / 100, math.Round(profit*100) / 100, err

}
