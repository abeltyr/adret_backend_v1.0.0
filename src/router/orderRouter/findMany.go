package orderRouter

import (
	"adr/backend/src/graphql/model"
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"adr/backend/src/service/orderService"
	"adr/backend/src/service/userService"
	"context"
	"encoding/json"
	"log"
)

func FindMany(ctx context.Context, input model.OrdersFilter) ([]*model.Order, error) {

	client := ctx.Value(srcModel.ConfigKey("client")).(*db.PrismaClient)
	clientCtx := ctx.Value(srcModel.ConfigKey("clientCtx")).(context.Context)
	cognitoUser := ctx.Value(srcModel.ConfigKey("user"))

	_, currentCompany, fetchErr := userService.Checking(cognitoUser, "", client, clientCtx)
	if fetchErr != nil {
		log.Println("error userService Checking", fetchErr)
		return nil, fetchErr
	}

	ctx = context.WithValue(ctx, srcModel.ConfigKey("currentCompany"), currentCompany)

	orderData, err := orderService.FindMany(input, client, ctx)
	if err != nil {
		log.Println("error productService FindMany", err)
		return nil, err
	}

	j, _ := json.Marshal(orderData)

	var order []*model.Order
	_ = json.Unmarshal(j, &order)

	return order, nil
}
