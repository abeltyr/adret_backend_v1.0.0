package userService

import (
	"adr/backend/src/graphql/model"
	srcModel "adr/backend/src/model"
	"adr/backend/src/utils"
	"context"
	"log"

	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func UpdatePersonalPassword(input model.UpdatePersonalPasswordInput, ctx context.Context) error {

	accessToken := ctx.Value(srcModel.ConfigKey("accessToken")).(string)

	_, err := utils.Cognito().ChangePassword(&cognito.ChangePasswordInput{
		PreviousPassword: input.OldPassword,
		ProposedPassword: input.Password,
		AccessToken:      &accessToken,
	})

	if err != nil {
		log.Println("user password update issue", err)
		return err
	}

	return nil

}
