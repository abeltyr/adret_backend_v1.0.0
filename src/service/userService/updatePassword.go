package userService

import (
	"adr/backend/src/graphql/model"
	"adr/backend/src/utils"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func UpdatePassword(input model.UpdateUserPasswordInput) error {

	userPool := os.Getenv("AWS_COGNITO_USER_POOL")

	_, err := utils.Cognito().AdminSetUserPassword(
		&cognito.AdminSetUserPasswordInput{
			Password:   aws.String(*input.Password),
			UserPoolId: aws.String(userPool),
			Username:   aws.String(*input.Username),
			Permanent:  aws.Bool(true),
		},
	)

	if err != nil {
		log.Println("cognito user password update error", err)
		return err
	}
	_, err = utils.Cognito().AdminUserGlobalSignOut(
		&cognito.AdminUserGlobalSignOutInput{
			UserPoolId: aws.String(userPool),
			Username:   aws.String(*input.Username),
		},
	)

	if err != nil {
		log.Println("cognito user sign out error", err)
		return err
	}

	return nil

}
