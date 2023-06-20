package utils

import (
	"adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"encoding/json"
	"errors"
	"fmt"
)

func CheckUser(user *db.UserModel, company *db.CompanyModel, role model.Role) error {

	userCompanyId, userCompanyIdOK := user.CompanyID()
	if !userCompanyIdOK {
		return errors.New("user doesn't have company")
	}

	if role == model.Owner {
		OwnerId, OwnerIdOk := company.OwnerID()
		if !OwnerIdOk {
			return errors.New("company doesn't have an owner")
		}
		if role == model.Owner && OwnerId != user.ID {
			return errors.New("user is not the owner of the company")
		}
	}

	if company.ID != userCompanyId {
		return errors.New("company id doesn't match")
	}

	if role == model.Manager && db.Role(role) != user.UserRole {
		return errors.New("user isn't a manger")
	}

	return nil
}

func GetCogitoUser(data interface{}) (model.CognitoUser, error) {
	currentUser := fmt.Sprintf("%v", data)

	var cognitoUser model.CognitoUser
	err := json.Unmarshal([]byte(currentUser), &cognitoUser)
	if err != nil {
		return model.CognitoUser{}, err
	}

	return cognitoUser, nil
}
