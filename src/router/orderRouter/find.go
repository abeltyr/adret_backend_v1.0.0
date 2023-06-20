package orderRouter

import (
	"adr/backend/src/graphql/model"
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"adr/backend/src/service/orderService"
	"adr/backend/src/utils"
	"context"
	"log"
)

func Find(ctx context.Context, id string) (*model.Order, error) {

	client := ctx.Value(srcModel.ConfigKey("client")).(*db.PrismaClient)

	orderData, err := orderService.Find(id, client, ctx)
	if err != nil {
		log.Println("error orderService Find", err)
		return nil, err
	}

	order := utils.OrderConverter(orderData)
	return order, nil
}
