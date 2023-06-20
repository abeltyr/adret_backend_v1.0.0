package inventoryRouter

import (
	"adr/backend/src/graphql/model"
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"adr/backend/src/service/inventoryService"
	"context"
	"encoding/json"
	"log"
)

func FindMany(ctx context.Context, input model.InventoriesFilter) ([]*model.Inventory, error) {

	client := ctx.Value(srcModel.ConfigKey("client")).(*db.PrismaClient)

	inventoryData, err := inventoryService.FindMany(input, client, ctx)
	if err != nil {
		log.Println("error productService FindMany", err)
		return nil, err
	}

	j, _ := json.Marshal(inventoryData)

	var inventories []*model.Inventory
	_ = json.Unmarshal(j, &inventories)

	return inventories, nil
}
