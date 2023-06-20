package userRouter

import (
	"adr/backend/src/graphql/model"
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"adr/backend/src/service/userService"
	"context"
	"log"
)

func UpdateAdminPassword(ctx context.Context, username string, password string) (bool, error) {

	client := ctx.Value(srcModel.ConfigKey("client")).(*db.PrismaClient)
	clientCtx := ctx.Value(srcModel.ConfigKey("clientCtx")).(context.Context)

	user, err := userService.FindByUsername(username, client, clientCtx)
	if user == nil || err != nil {
		log.Println("error userService FindBySub", err)
		return false, err
	}

	err = userService.UpdatePassword(model.UpdateUserPasswordInput{
		Username: &username,
		Password: &password,
	})
	if err != nil {
		return false, err
	}
	return true, err
}
