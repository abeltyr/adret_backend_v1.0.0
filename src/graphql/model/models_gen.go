// Code generated by github.com/arsmn/fastgql, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"

	"github.com/arsmn/fastgql/graphql"
)

type Company struct {
	ID          string  `json:"id"`
	Name        *string `json:"name"`
	CompanyCode *string `json:"companyCode"`
	Detail      *string `json:"detail"`
	Longitude   *string `json:"longitude"`
	Latitude    *bool   `json:"latitude"`
	OwnerID     *string `json:"ownerId"`
	Owner       *User   `json:"owner"`
	CreatedAt   *string `json:"createdAt"`
	UpdatedAt   *string `json:"updatedAt"`
	DeletedAt   *string `json:"deletedAt"`
}

type File struct {
	Name       *string `json:"name"`
	URL        *string `json:"url"`
	PreviewURL *string `json:"previewUrl"`
	Width      *string `json:"width"`
	CreatorID  *string `json:"creatorId"`
	Height     *User   `json:"height"`
	Size       *string `json:"size"`
}

type FilterInput struct {
	After  *string `json:"after"`
	Before *string `json:"before"`
	Limit  *int    `json:"limit"`
}

type Inventory struct {
	ID                        string                `json:"id"`
	InventoryVariation        []*InventoryVariation `json:"inventoryVariation"`
	Media                     *string               `json:"media"`
	SalesAmount               *int                  `json:"salesAmount"`
	Available                 *int                  `json:"available"`
	InitialPrice              *float64              `json:"initialPrice"`
	MinSellingPriceEstimation *float64              `json:"minSellingPriceEstimation"`
	MaxSellingPriceEstimation *float64              `json:"maxSellingPriceEstimation"`
	ProductID                 *string               `json:"productId"`
	Product                   *Product              `json:"product"`
	CreatedAt                 *string               `json:"createdAt"`
	UpdatedAt                 *string               `json:"updatedAt"`
	DeletedAt                 *string               `json:"deletedAt"`
}

type InventoryVariation struct {
	ID                 *string `json:"id"`
	Data               *string `json:"data"`
	Title              *string `json:"title"`
	ProductVariationID *string `json:"productVariationId"`
}

type OnlineOrderDetail struct {
	ID             string  `json:"id"`
	Delivery       *bool   `json:"delivery"`
	Latitude       *string `json:"latitude"`
	Longitude      *string `json:"longitude"`
	LocationName   *string `json:"locationName"`
	Detail         *string `json:"detail"`
	PhoneNumber    *string `json:"phoneNumber"`
	Email          *string `json:"email"`
	ClientFullName *string `json:"clientFullName"`
	RedirectURL    *string `json:"redirectUrl"`
	Status         *string `json:"status"`
	Extra          *string `json:"extra"`
}

type OnlineOrderPayment struct {
	ID        string  `json:"id"`
	Method    *string `json:"method"`
	Reference *string `json:"reference"`
	TxRef     *string `json:"txRef"`
	Paid      *bool   `json:"paid"`
	PaidDate  *string `json:"paidDate"`
	Extra     *string `json:"extra"`
}

type Order struct {
	ID                 string              `json:"id"`
	Online             *bool               `json:"online"`
	OrderNumber        *string             `json:"orderNumber"`
	Note               *string             `json:"note"`
	TotalPrice         *float64            `json:"totalPrice"`
	TotalProfit        *float64            `json:"totalProfit"`
	Date               *string             `json:"date"`
	SellerID           *string             `json:"sellerId"`
	Seller             *User               `json:"seller"`
	CompanyID          *string             `json:"companyId"`
	Company            *Company            `json:"company"`
	CreatedAt          *string             `json:"createdAt"`
	UpdatedAt          *string             `json:"updatedAt"`
	DeletedAt          *string             `json:"deletedAt"`
	OnlineOrderDetail  *OnlineOrderDetail  `json:"onlineOrderDetail"`
	OnlineOrderPayment *OnlineOrderPayment `json:"onlineOrderPayment"`
	Sales              []*Sales            `json:"sales"`
}

type Product struct {
	ID          string              `json:"id"`
	ProductCode *string             `json:"productCode"`
	Title       *string             `json:"title"`
	Detail      *string             `json:"detail"`
	CreatorID   *string             `json:"creatorId"`
	Creator     *User               `json:"creator"`
	CompanyID   *string             `json:"companyId"`
	Company     *Company            `json:"company"`
	Category    *string             `json:"category"`
	InStock     *int                `json:"inStock"`
	Media       []*string           `json:"media"`
	Inventory   []*Inventory        `json:"inventory"`
	Variation   []*ProductVariation `json:"variation"`
	CreatedAt   *string             `json:"createdAt"`
	UpdatedAt   *string             `json:"updatedAt"`
	DeletedAt   *string             `json:"deletedAt"`
}

type ProductVariation struct {
	ID    *string `json:"id"`
	Title *string `json:"title"`
	Order *int    `json:"order"`
}

type Sales struct {
	ID           string     `json:"id"`
	SellingPrice *float64   `json:"sellingPrice"`
	Profit       *float64   `json:"profit"`
	InventoryID  *string    `json:"inventoryId"`
	Inventory    *Inventory `json:"inventory"`
	OrderID      *string    `json:"orderId"`
	Order        *Order     `json:"order"`
}

type Summary struct {
	ID               string       `json:"id"`
	Earning          *float64     `json:"earning"`
	Profit           *float64     `json:"profit"`
	ManagerAccepted  *bool        `json:"managerAccepted"`
	ManagerID        *string      `json:"managerId"`
	Manager          *User        `json:"manager"`
	EmployeeID       *string      `json:"employeeId"`
	Employee         *User        `json:"employee"`
	Date             *string      `json:"date"`
	StartDate        *string      `json:"startDate"`
	EndDate          *string      `json:"endDate"`
	SummaryInventory []*Inventory `json:"summaryInventory"`
	CreatedAt        *string      `json:"createdAt"`
	UpdatedAt        *string      `json:"updatedAt"`
	DeletedAt        *string      `json:"deletedAt"`
}

type User struct {
	ID          string   `json:"id"`
	FullName    *string  `json:"fullName"`
	PhoneNumber *string  `json:"phoneNumber"`
	Email       *string  `json:"email"`
	UserName    *string  `json:"userName"`
	IsActive    *bool    `json:"isActive"`
	UserRole    *Role    `json:"userRole"`
	CompanyID   *string  `json:"companyId"`
	Company     *Company `json:"company"`
	CreatorID   *string  `json:"creatorId"`
	Creator     *User    `json:"creator"`
	CreatedAt   *string  `json:"createdAt"`
	UpdatedAt   *string  `json:"updatedAt"`
	DeletedAt   *string  `json:"deletedAt"`
}

type CreateBranchInput struct {
	BranchName string `json:"branchName"`
	Latitude   string `json:"latitude"`
	Longitude  string `json:"longitude"`
}

type CreateCompanyInput struct {
	Name        string `json:"name"`
	CompanyCode string `json:"companyCode"`
}

type CreateInventoryInput struct {
	Amount                    int                      `json:"amount"`
	InitialPrice              float64                  `json:"initialPrice"`
	MinSellingPriceEstimation float64                  `json:"minSellingPriceEstimation"`
	MaxSellingPriceEstimation float64                  `json:"maxSellingPriceEstimation"`
	Media                     *int                     `json:"media"`
	Variation                 []*ProductVariationInput `json:"variation"`
}

type CreateLocalOrderInput struct {
	Note       *string       `json:"note"`
	SalesInput []*SalesInput `json:"salesInput"`
}

type CreateOwnerInput struct {
	PhoneNumber string `json:"phoneNumber"`
	FullName    string `json:"fullName"`
	UserName    string `json:"userName"`
	Password    string `json:"password"`
}

type CreateProductInput struct {
	Title     string                  `json:"title"`
	Detail    string                  `json:"detail"`
	Category  string                  `json:"category"`
	Media     []*graphql.Upload       `json:"media"`
	Inventory []*CreateInventoryInput `json:"inventory"`
}

type CreateUserInput struct {
	PhoneNumber *string `json:"phoneNumber"`
	FullName    *string `json:"fullName"`
	UserName    *string `json:"userName"`
	Password    *string `json:"password"`
}

type EmployeeDailySummaryFilter struct {
	EmployeeID string `json:"employeeId"`
	Date       string `json:"date"`
}

type InventoriesFilter struct {
	Filter      *FilterInput `json:"filter"`
	Sold        *bool        `json:"sold"`
	BoughtDates *string      `json:"boughtDates"`
	ProductID   *string      `json:"productId"`
	EmployeeID  *string      `json:"employeeId"`
}

type ManagerAcceptInput struct {
	EmployeeID string `json:"employeeId"`
	Date       string `json:"date"`
}

type OrdersFilter struct {
	Filter        *FilterInput `json:"filter"`
	SellerID      *string      `json:"sellerId"`
	MinTotalPrice *float64     `json:"minTotalPrice"`
	MaxTotalPrice *float64     `json:"maxTotalPrice"`
	StartDate     *string      `json:"startDate"`
	EndDate       *string      `json:"endDate"`
}

type ProductVariationInput struct {
	Title *string `json:"title"`
	Data  *string `json:"data"`
}

type ProductsFilter struct {
	Filter     *FilterInput `json:"filter"`
	Code       *string      `json:"code"`
	Title      *string      `json:"title"`
	Category   *string      `json:"category"`
	CreatorID  *string      `json:"creatorId"`
	TopSelling *bool        `json:"topSelling"`
}

type SalesFilter struct {
	Filter          *FilterInput `json:"filter"`
	OrderID         *string      `json:"orderId"`
	InventoryID     *string      `json:"inventoryId"`
	MinSellingPrice *float64     `json:"minSellingPrice"`
	MaxSellingPrice *float64     `json:"maxSellingPrice"`
}

type SalesInput struct {
	SellingPrice float64 `json:"sellingPrice"`
	InventoryID  string  `json:"inventoryId"`
	Amount       int     `json:"amount"`
}

type UpdateInventoryInput struct {
	ProductID                 string   `json:"productId"`
	Title                     string   `json:"title"`
	NewTitle                  *string  `json:"newTitle"`
	InitialPrice              *float64 `json:"initialPrice"`
	MinSellingPriceEstimation *float64 `json:"minSellingPriceEstimation"`
	MaxSellingPriceEstimation *float64 `json:"maxSellingPriceEstimation"`
}

type UpdateMedia struct {
	File *graphql.Upload `json:"file"`
	URL  *string         `json:"url"`
}

type UpdatePersonalPasswordInput struct {
	Password    *string `json:"password"`
	OldPassword *string `json:"oldPassword"`
}

type UpdateProductInput struct {
	Title     *string                        `json:"title"`
	Detail    *string                        `json:"detail"`
	Category  *string                        `json:"category"`
	Media     []*UpdateMedia                 `json:"media"`
	Inventory []*UpdateProductInventoryInput `json:"inventory"`
	ID        string                         `json:"id"`
}

type UpdateProductInventoryInput struct {
	Amount                    *int                     `json:"amount"`
	InitialPrice              *float64                 `json:"initialPrice"`
	MinSellingPriceEstimation *float64                 `json:"minSellingPriceEstimation"`
	MaxSellingPriceEstimation *float64                 `json:"maxSellingPriceEstimation"`
	Media                     *int                     `json:"media"`
	Variation                 []*ProductVariationInput `json:"variation"`
	ID                        *string                  `json:"id"`
}

type UpdateUserInput struct {
	PhoneNumber *string `json:"phoneNumber"`
	FullName    *string `json:"fullName"`
	Email       *string `json:"email"`
	ID          *string `json:"id"`
}

type UpdateUserPasswordInput struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}

type UsersFilter struct {
	Filter    *FilterInput `json:"filter"`
	CompanyID *string      `json:"companyId"`
	Role      *Role        `json:"role"`
	ExceptMe  *bool        `json:"exceptMe"`
}

type Role string

const (
	RoleManager  Role = "Manager"
	RoleEmployee Role = "Employee"
)

var AllRole = []Role{
	RoleManager,
	RoleEmployee,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleManager, RoleEmployee:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
