package sharedRouter

import (
	"adr/backend/src/graphql/model"
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"adr/backend/src/service/inventoryService"
	"adr/backend/src/utils"
	"context"
	"log"
)

func Inventory(ctx context.Context, id string) (*model.Inventory, error) {

	client := ctx.Value(srcModel.ConfigKey("client")).(*db.PrismaClient)

	data, err := inventoryService.Find(id, client, ctx)
	if err != nil {
		log.Println("error inventoryService Find sharedRouter", err)
		return nil, err
	}

	inventory := utils.InventoryConverter(data)

	return inventory, nil
}
