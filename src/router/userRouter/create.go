package userRouter

import (
	"adr/backend/src/graphql/model"
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"adr/backend/src/service/userService"
	"adr/backend/src/utils"
	"context"
	"errors"
	"log"
)

func Create(ctx context.Context, input model.CreateUserInput) (*model.User, error) {

	client := ctx.Value(srcModel.ConfigKey("client")).(*db.PrismaClient)
	clientCtx := ctx.Value(srcModel.ConfigKey("clientCtx")).(context.Context)
	cognitoUser := ctx.Value(srcModel.ConfigKey("user"))

	// fetch the requesting user
	currentUser, currentCompany, err := userService.Checking(cognitoUser, srcModel.Owner, client, clientCtx)
	if err != nil {
		log.Println("error userService Checking", err)
		return nil, err
	}

	if currentCompany == nil {
		return nil, errors.New("the creating user needs a company ")
	}

	ctx = context.WithValue(ctx, srcModel.ConfigKey("currentUser"), currentUser)
	exceptMe := true
	users, err := userService.FindMany(model.UsersFilter{
		CompanyID: &currentCompany.ID,
		ExceptMe:  &exceptMe,
	}, client, ctx)

	if err != nil {
		log.Println("error userService FindMany", err)
		return nil, err
	}

	if len(users) > 3 {
		return nil, errors.New("a company can have max 3 employees at the moment")
	}

	userData, err := userService.Create(input, srcModel.Employee, currentCompany, currentUser, client, clientCtx)
	if err != nil {
		log.Println("error userService Create", err)
		return nil, err
	}

	user := utils.UserConverter(userData)
	return user, nil
}
