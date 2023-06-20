package userService

import (
	graphModel "adr/backend/src/graphql/model"
	"adr/backend/src/prisma/db"
	"adr/backend/src/utils"
	"context"
	"errors"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func Update(input graphModel.UpdateUserInput, client *db.PrismaClient, ctx context.Context) (*db.UserModel, error) {

	user, err := Find(*input.ID, client, ctx)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("invalid id")
	}

	email, _ := user.Email()
	username := user.UserName
	phoneNumber, _ := user.PhoneNumber()
	fullName := user.FullName

	parameter := []db.UserSetParam{}

	if input.PhoneNumber != nil {
		if *input.PhoneNumber != phoneNumber {
			parameter = append(parameter,
				db.User.PhoneNumber.Set(*input.PhoneNumber),
			)
		}
		phoneNumber = *input.PhoneNumber
	}

	if input.FullName != nil {
		if *input.FullName != fullName {
			parameter = append(parameter,
				db.User.FullName.Set(*input.FullName),
			)
		}
		fullName = *input.FullName
	}

	if input.Email != nil {
		if *input.Email != email {
			parameter = append(parameter,
				db.User.Email.Set(*input.Email),
			)
		}
		email = *input.Email
	}

	userPool := os.Getenv("AWS_COGNITO_USER_POOL")

	_, err = utils.Cognito().AdminUpdateUserAttributes(&cognito.AdminUpdateUserAttributesInput{
		Username:   aws.String(username),
		UserPoolId: aws.String(userPool),
		UserAttributes: []*cognito.AttributeType{
			{
				Name:  aws.String("phone_number"),
				Value: aws.String(phoneNumber),
			},
			{
				Name:  aws.String("given_name"),
				Value: aws.String(fullName),
			},
			{
				Name:  aws.String("email"),
				Value: aws.String(email),
			},
		},
	})

	if err != nil {
		log.Println("Cognito user update error", err)
		return nil, err
	}

	updated, err := client.User.
		FindUnique(
			db.User.ID.Equals(*input.ID),
		).
		Update(parameter[:]...).
		Exec(ctx)
	if err != nil {
		log.Println("user update error", err)
		return nil, err
	}

	return updated, err
}
