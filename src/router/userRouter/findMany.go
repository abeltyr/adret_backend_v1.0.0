package userRouter

import (
	"adr/backend/src/graphql/model"
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"adr/backend/src/service/userService"
	"context"
	"encoding/json"
	"log"
)

func FindMany(ctx context.Context, input model.UsersFilter) ([]*model.User, error) {

	client := ctx.Value(srcModel.ConfigKey("client")).(*db.PrismaClient)
	clientCtx := ctx.Value(srcModel.ConfigKey("clientCtx")).(context.Context)
	cognitoUser := ctx.Value(srcModel.ConfigKey("user"))

	currentUser, company, fetchErr := userService.Checking(cognitoUser, srcModel.Manager, client, clientCtx)
	if fetchErr != nil {
		log.Println("error userService Checking", fetchErr)
		return nil, fetchErr
	}

	ctx = context.WithValue(ctx, srcModel.ConfigKey("currentUser"), currentUser)
	input.CompanyID = &company.ID
	userData, err := userService.FindMany(input, client, ctx)
	if err != nil {
		log.Println("error userService FindMany", err)
		return nil, err
	}

	j, _ := json.Marshal(userData)

	var users []*model.User
	_ = json.Unmarshal(j, &users)

	return users, nil
}
