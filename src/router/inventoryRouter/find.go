package inventoryRouter

import (
	"adr/backend/src/graphql/model"
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"adr/backend/src/service/inventoryService"
	"adr/backend/src/utils"
	"context"
	"log"
)

func Find(ctx context.Context, id string) (*model.Inventory, error) {

	client := ctx.Value(srcModel.ConfigKey("client")).(*db.PrismaClient)

	inventoryData, err := inventoryService.Find(id, client, ctx)
	if err != nil {
		log.Println("error inventoryService Find", err)
		return nil, err
	}

	inventory := utils.InventoryConverter(inventoryData)
	return inventory, nil
}
