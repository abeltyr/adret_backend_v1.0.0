package summaryRouter

import (
	"adr/backend/src/graphql/model"
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"adr/backend/src/service/inventoryService"
	"context"
	"encoding/json"
	"log"
)

func InventoryList(ctx context.Context, obj *model.Summary, filter *model.FilterInput) ([]*model.Inventory, error) {

	j, _ := json.Marshal(obj)

	var summaryData *db.EmployeeDailySummaryModel
	_ = json.Unmarshal(j, &summaryData)

	client := ctx.Value(srcModel.ConfigKey("client")).(*db.PrismaClient)

	sold := true
	employeeId := summaryData.EmployeeID

	filterData := model.InventoriesFilter{
		Filter:      filter,
		Sold:        &sold,
		BoughtDates: nil,
		ProductID:   nil,
		EmployeeID:  &employeeId,
	}

	inventoryData, err := inventoryService.FindMany(filterData, client, ctx)
	if err != nil {
		log.Println("error inventoryService FindMany", err)
		return nil, err
	}

	object, _ := json.Marshal(inventoryData)

	var inventories []*model.Inventory
	_ = json.Unmarshal(object, &inventories)

	return inventories, nil
}
