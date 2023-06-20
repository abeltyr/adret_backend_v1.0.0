package companyBranchService

type CompanyBranchInput struct {
	BranchName string `json:"branchName"`
	Latitude   string `json:"latitude"`
	Longitude  string `json:"longitude"`
	CompanyId  string `json:"companyId"`
}
