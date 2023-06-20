package salesRouter

import (
	"adr/backend/src/graphql/model"
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"adr/backend/src/service/saleService"
	"adr/backend/src/utils"
	"context"
	"log"
)

func Find(ctx context.Context, orderID string, InventoryID string) (*model.Sales, error) {

	client := ctx.Value(srcModel.ConfigKey("client")).(*db.PrismaClient)

	saleData, err := saleService.Find(orderID, InventoryID, client, ctx)
	if err != nil {
		log.Println("error saleService Find", err)
		return nil, err
	}

	sale := utils.SalesConverter(saleData)
	return sale, nil
}
