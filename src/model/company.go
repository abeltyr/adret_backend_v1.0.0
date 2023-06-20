package model

type Company struct {
	Id          string `json:"id"`
	OwnerId     string `json:"ownerId"`
	Name        string `json:"name"`
	CompanyCode string `json:"companyCode"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type CompanyCreateInput struct {
	CompanyData Company `json:"companyData"`
}

type CompanyUpdateInput struct {
	CompanyData Company `json:"companyData"`
	Id          IdFetch `json:"id"`
}
