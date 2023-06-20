package userService

import (
	"adr/backend/src/graphql/model"
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"adr/backend/src/utils"
	"context"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func Create(input model.CreateUserInput, role srcModel.Role, company *db.CompanyModel, user *db.UserModel, client *db.PrismaClient, ctx context.Context) (*db.UserModel, error) {

	userPool := os.Getenv("AWS_COGNITO_USER_POOL")

	userName := company.CompanyCode + "-" + strings.ToLower(*input.UserName)
	data, err := utils.Cognito().AdminCreateUser(&cognito.AdminCreateUserInput{
		Username:   aws.String(userName),
		UserPoolId: aws.String(userPool),
		UserAttributes: []*cognito.AttributeType{
			{
				Name:  aws.String("phone_number"),
				Value: aws.String(*input.PhoneNumber),
			},
			{
				Name:  aws.String("given_name"),
				Value: aws.String(*input.FullName),
			},
		},
	})

	if err != nil {
		log.Println("user cognito creation err for ", input.UserName, "error:", err)
		return nil, err
	}

	err = UpdatePassword(model.UpdateUserPasswordInput{
		Username: &userName,
		Password: input.Password,
	})

	if err != nil {
		utils.Cognito().AdminDeleteUser(&cognitoidentityprovider.AdminDeleteUserInput{
			UserPoolId: &userPool,
			Username:   &userName,
		})
		log.Println("user password update err", input.UserName, "error:", err)
		return nil, err
	}

	parameter := []db.UserSetParam{
		db.User.UserRole.Set(db.Role(role)),

		db.User.CompanyID.Set(company.ID),
	}

	if data.User.Attributes[0].Value != nil {
		parameter = append(parameter, db.User.SubID.Set(*data.User.Attributes[0].Value))
	}
	if input.PhoneNumber != nil {
		parameter = append(parameter, db.User.PhoneNumber.Set(*input.PhoneNumber))
	}
	if user != nil {
		branchId, ok := user.BranchID()
		if ok && branchId != "" {
			parameter = append(parameter, db.User.BranchID.Set(branchId))
		}
	}

	if user != nil {
		parameter = append(parameter, db.User.CreatorID.Set(user.ID))
	} else {
		parameter = append(parameter, db.User.UserRole.Set(db.RoleManager))
	}

	createdUser, err := client.User.CreateOne(
		db.User.FullName.Set(*input.FullName),
		db.User.UserName.Set(userName),
		parameter[:]...,
	).Exec(ctx)

	if err != nil {
		utils.Cognito().AdminDeleteUser(&cognitoidentityprovider.AdminDeleteUserInput{
			UserPoolId: &userPool,
			Username:   &userName,
		})
		log.Println("user creation database side error", input.UserName, "error:", err)
		return nil, err
	}

	return createdUser, nil

}
