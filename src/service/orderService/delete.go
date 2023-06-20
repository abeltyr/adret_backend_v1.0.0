package orderService

import (
	"adr/backend/src/prisma/db"
	"context"
)

func Delete(id string, client *db.PrismaClient, ctx context.Context) (*db.OrderModel, error) {

	createdOrder, err := client.Order.FindUnique(db.Order.ID.Equals(id)).Delete().Exec(ctx)

	if err != nil {
		return nil, err
	}

	return createdOrder, nil

}
