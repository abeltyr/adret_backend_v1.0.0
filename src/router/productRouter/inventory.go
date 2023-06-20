package productRouter

import (
	"adr/backend/src/graphql/model"
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"adr/backend/src/service/inventoryService"
	"context"
	"encoding/json"
	"log"
)

func Inventory(ctx context.Context, obj *model.Product) ([]*model.Inventory, error) {

	client := ctx.Value(srcModel.ConfigKey("client")).(*db.PrismaClient)
	data := model.InventoriesFilter{
		ProductID: &obj.ID,
	}
	inventoryData, err := inventoryService.FindMany(
		data, client, ctx)
	if err != nil {
		log.Println("error  inventoryService FindMany", err)
		return nil, err
	}

	j, _ := json.Marshal(inventoryData)

	var inventories []*model.Inventory
	_ = json.Unmarshal(j, &inventories)

	return inventories, nil
}
