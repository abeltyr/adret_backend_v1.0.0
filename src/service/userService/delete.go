package userService

import (
	"adr/backend/src/prisma/db"
	"adr/backend/src/utils"
	"context"
	"os"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func Delete(id string, client *db.PrismaClient, ctx context.Context) error {

	user, err := Find(id, client, ctx)
	if err != nil {
		return err
	}

	if user != nil {
		userPool := os.Getenv("AWS_COGNITO_USER_POOL")
		_, err = utils.Cognito().AdminDeleteUser(&cognitoidentityprovider.AdminDeleteUserInput{
			UserPoolId: &userPool,
			Username:   &user.UserName,
		})
		if err != nil {
			return err
		}

		_, err = client.User.FindUnique(
			db.User.ID.Equals(id),
		).Delete().Exec(ctx)

		if err != nil {
			return err
		}
	}

	return nil

}
