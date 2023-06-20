package model

type ProductCreateInput struct {
	Title       string                 `json:"title"`
	Detail      string                 `json:"detail"`
	Category    string                 `json:"category"`
	CompanyCode string                 `json:"companyCode"`
	CompanyId   string                 `json:"companyId"`
	CreatorId   string                 `json:"creatorId"`
	Inventory   []InventoryCreateInput `json:"inventory"`
}

type InventoryCreateInput struct {
	Title                  string  `json:"title"`
	Amount                 uint16  `json:"amount"`
	ProductId              string  `json:"productId"`
	InitialPrice           float64 `json:"initialPrice"`
	SellingPriceEstimation float64 `json:"sellingPriceEstimation"`
}

type ProductUpdateInput struct {
	Title  string `json:"title"`
	Detail string `json:"detail"`
	Id     string `json:"id"`
}

type ProductFindManyFilter struct {
	Filter    FilterData `json:"filter"`
	CompanyId string     `json:"companyId"`
	CreatorId string     `json:"creatorId"`
	Sold      bool       `json:"sold"`
}

type ProductIFindFilter struct {
	InventoryFilter FilterData `json:"inventoryFilter"`
	Id              string     `json:"id"`
	Sold            bool       `json:"sold"`
}
