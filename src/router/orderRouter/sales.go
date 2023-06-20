package orderRouter

import (
	"adr/backend/src/graphql/model"
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"adr/backend/src/service/saleService"
	"context"
	"encoding/json"
	"log"
)

func Sales(ctx context.Context, orderID string) ([]*model.Sales, error) {
	client := ctx.Value(srcModel.ConfigKey("client")).(*db.PrismaClient)
	data, err := saleService.FindMany(model.SalesFilter{
		OrderID: &orderID,
	}, client, ctx)
	if err != nil {
		log.Println("error saleService FindMany", err)
		return nil, err
	}

	j, _ := json.Marshal(data)

	var sales []*model.Sales
	_ = json.Unmarshal(j, &sales)

	return sales, nil
}
