package userRouter

import (
	"adr/backend/src/graphql/model"
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"adr/backend/src/service/userService"
	"context"
	"errors"
	"log"
)

func UpdatePassword(ctx context.Context, input model.UpdateUserPasswordInput) (*bool, error) {

	client := ctx.Value(srcModel.ConfigKey("client")).(*db.PrismaClient)
	clientCtx := ctx.Value(srcModel.ConfigKey("clientCtx")).(context.Context)
	cognitoUser := ctx.Value(srcModel.ConfigKey("user"))

	_, company, fetchErr := userService.Checking(cognitoUser, srcModel.Manager, client, clientCtx)
	if fetchErr != nil {
		log.Println("error userService Checking", fetchErr)
		return nil, fetchErr
	}

	user, err := userService.FindByUsername(*input.Username, client, clientCtx)
	if user == nil || err != nil {
		log.Println("error userService FindBySub", err)
		return nil, err
	}
	CompanyID, ok := user.CompanyID()
	if !ok || (ok && company.ID != CompanyID) {
		err = errors.New("user not a member of your shop")
		log.Println("error company id doesn't match", err)
		return nil, err
	}
	err = userService.UpdatePassword(input)

	done := true

	return &done, err
}
