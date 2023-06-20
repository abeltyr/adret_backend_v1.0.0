package productRouter

import (
	"adr/backend/src/graphql/model"
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"adr/backend/src/service/inventoryService"
	"context"
	"log"
)

func Remaining(ctx context.Context, productId string) (*int, error) {

	client := ctx.Value(srcModel.ConfigKey("client")).(*db.PrismaClient)

	inventories, error := inventoryService.FindMany(model.InventoriesFilter{
		ProductID: &productId,
	},
		client,
		ctx,
	)
	if error != nil {
		log.Println("error inventoryService FindMany", error)
		return nil, error
	}

	amount := 0
	sold := 0

	for _, inventory := range inventories {
		amount = amount + inventory.Available
		sold = sold + inventory.SalesAmount
	}

	remaining := amount - sold

	return &remaining, nil
}
