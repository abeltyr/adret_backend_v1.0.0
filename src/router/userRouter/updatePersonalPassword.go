package userRouter

import (
	"adr/backend/src/graphql/model"
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"adr/backend/src/service/userService"
	"context"
	"log"
)

func UpdatePersonalPassword(ctx context.Context, input model.UpdatePersonalPasswordInput) (*bool, error) {

	client := ctx.Value(srcModel.ConfigKey("client")).(*db.PrismaClient)
	clientCtx := ctx.Value(srcModel.ConfigKey("clientCtx")).(context.Context)
	cognitoUser := ctx.Value(srcModel.ConfigKey("user"))

	_, _, fetchErr := userService.Checking(cognitoUser, srcModel.Manager, client, clientCtx)
	if fetchErr != nil {
		log.Println("error userService Checking", fetchErr)
		return nil, fetchErr
	}

	err := userService.UpdatePersonalPassword(input, ctx)

	done := true

	return &done, err
}
