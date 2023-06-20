package userRouter

import (
	"adr/backend/src/graphql/model"
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"adr/backend/src/service/userService"
	"adr/backend/src/utils"
	"context"
	"log"
)

func Update(ctx context.Context, input model.UpdateUserInput) (*model.User, error) {

	client := ctx.Value(srcModel.ConfigKey("client")).(*db.PrismaClient)
	clientCtx := ctx.Value(srcModel.ConfigKey("clientCtx")).(context.Context)
	cognitoUser := ctx.Value(srcModel.ConfigKey("user"))

	currentUser, _, fetchErr := userService.Checking(cognitoUser, "", client, clientCtx)
	if fetchErr != nil {
		log.Println("error userService Checking", fetchErr)
		return nil, fetchErr
	}

	if currentUser.UserRole == db.Role(srcModel.Employee) {
		input.ID = &currentUser.ID
	}

	userData, err := userService.Update(input, client, ctx)
	if err != nil {
		log.Println("error userService Update", err)
		return nil, err
	}

	user := utils.UserConverter(userData)
	return user, nil
}
