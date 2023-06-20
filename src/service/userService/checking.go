package userService

import (
	"adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"adr/backend/src/service/companyService"
	"adr/backend/src/utils"
	"context"
	"errors"
)

func Checking(data interface{}, role model.Role, client *db.PrismaClient, ctx context.Context) (*db.UserModel, *db.CompanyModel, error) {

	// fetch the requesting user
	cognitoUser, fetchErr := utils.GetCogitoUser(data)
	if fetchErr != nil {
		return nil, nil, fetchErr
	}

	currentUser, err := FindBySub(cognitoUser.Sub, client, ctx)
	if err != nil {
		return nil, nil, err
	}

	//fetch the company id from the user
	companyId, ok := currentUser.CompanyID()
	if !ok {
		return nil, nil, errors.New("user doesn't have company")
	}
	company, companyErr := companyService.Find(companyId, client, ctx)
	if companyErr != nil {
		return nil, nil, companyErr
	}

	if role != "" {
		// check the user role to be owner
		CheckErr := utils.CheckUser(currentUser, company, role)
		if CheckErr != nil {
			return nil, nil, CheckErr
		}
	}

	return currentUser, company, nil
}
