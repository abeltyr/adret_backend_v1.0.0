package model

type Category struct {
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type CategoryProduct struct {
	CategoryName string `json:"categoryName"`
	ProductId    string `json:"productId"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
}

type CategoryCompany struct {
	CategoryName string `json:"categoryName"`
	CompanyId    string `json:"companyId"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
}

type CategoryProductInput struct {
	CategoryName string `json:"categoryName"`
	ProductId    string `json:"productId"`
}

type CategoryUpdateInput struct {
	CurrentName string `json:"currentName"`
	Name        string `json:"name"`
}

type CategoryCompanyInput struct {
	CategoryName string `json:"categoryName"`
	CompanyId    string `json:"companyId"`
}

type CategoryProductFindManyFilter struct {
	Filter          FilterData           `json:"filter"`
	CategoryProduct CategoryProductInput `json:"categoryProduct"`
}

type CategoryCompanyFindManyFilter struct {
	Filter          FilterData           `json:"filter"`
	CategoryCompany CategoryCompanyInput `json:"categoryCompany"`
}
