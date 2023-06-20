package variationRouter

import (
	"adr/backend/src/graphql/model"
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"adr/backend/src/service/inventoryVariationService"
	"context"
	"encoding/json"
	"log"
)

func FindMany(ctx context.Context, id string) ([]*model.InventoryVariation, error) {

	client := ctx.Value(srcModel.ConfigKey("client")).(*db.PrismaClient)

	inventoryVariationData, err := inventoryVariationService.FindMany(id, client, ctx)
	if err != nil {
		log.Println("error inventoryService FindMany", err)
		return nil, err
	}

	j, _ := json.Marshal(inventoryVariationData)

	var inventoryVariation []*model.InventoryVariation
	_ = json.Unmarshal(j, &inventoryVariation)

	return inventoryVariation, nil
}
