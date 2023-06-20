package model

type CognitoUser struct {
	Sub        string `json:"sub"`
	Iss        string `json:"iss"`
	Client_id  string `json:"client_id"`
	Origin_jti string `json:"origin_jti"`
	Event_id   string `json:"event_id"`
	Token_use  string `json:"token_use"`
	Scope      string `json:"scope"`
	Auth_time  int64  `json:"auth_time"`
	Exp        int64  `json:"exp"`
	Iat        int64  `json:"iat"`
	Jti        string `json:"jti"`
	Username   string `json:"username"`
}

type UserFindManyFilter struct {
	Filter    FilterData `json:"filter"`
	CompanyId string     `json:"companyId"`
	Role      string     `json:"role"`
}

type Role string

const (
	Owner    Role = "Owner"
	Manager  Role = "Manager"
	Employee Role = "Employee"
)
