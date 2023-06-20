package orderService

import (
	"adr/backend/src/prisma/db"
	"context"
)

func Find(id string, client *db.PrismaClient, ctx context.Context) (*db.OrderModel, error) {

	sales, err := client.Order.FindUnique(
		db.Order.ID.Equals(id),
	).Exec(ctx)

	return sales, err
}
