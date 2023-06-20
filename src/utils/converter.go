package utils

import (
	"adr/backend/src/graphql/model"
	"encoding/json"
)

func LimitSetter(limit int32) int32 {
	if limit == 0 || limit >= 50 {
		limit = 50
	}
	return limit
}

func UserConverter(userData interface{}) *model.User {
	j, _ := json.Marshal(userData)

	var user *model.User
	_ = json.Unmarshal(j, &user)

	return user
}

func ProductConverter(data interface{}) *model.Product {
	j, _ := json.Marshal(data)

	var product *model.Product
	_ = json.Unmarshal(j, &product)

	return product
}

func CompanyConverter(data interface{}) *model.Company {
	j, _ := json.Marshal(data)

	var company *model.Company
	_ = json.Unmarshal(j, &company)

	return company
}

func InventoryConverter(data interface{}) *model.Inventory {
	j, _ := json.Marshal(data)

	var inventory *model.Inventory
	_ = json.Unmarshal(j, &inventory)

	return inventory
}

func SummaryConverter(data interface{}) *model.Summary {
	j, _ := json.Marshal(data)

	var summary *model.Summary
	_ = json.Unmarshal(j, &summary)

	return summary
}

func InventoryVariationConverter(data interface{}) *model.InventoryVariation {
	j, _ := json.Marshal(data)

	var inventoryVariation *model.InventoryVariation
	_ = json.Unmarshal(j, &inventoryVariation)

	return inventoryVariation
}

func SalesConverter(data interface{}) *model.Sales {
	j, _ := json.Marshal(data)

	var sales *model.Sales
	_ = json.Unmarshal(j, &sales)

	return sales
}

func OrderConverter(data interface{}) *model.Order {
	j, _ := json.Marshal(data)

	var order *model.Order
	_ = json.Unmarshal(j, &order)

	return order
}
