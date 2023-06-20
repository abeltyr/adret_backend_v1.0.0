package orderService

import (
	"adr/backend/src/prisma/db"
	"context"
)

func Update(id string, totalPrice float64, totalProfit float64, client *db.PrismaClient, ctx context.Context) (*db.OrderModel, error) {

	createdSale, err := client.Order.FindUnique(
		db.Order.ID.Equals(id),
	).Update(
		db.Order.TotalPrice.Set(totalPrice),
		db.Order.TotalProfit.Set(totalProfit),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return createdSale, nil

}
