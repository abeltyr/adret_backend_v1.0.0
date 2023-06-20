package sharedRouter

import (
	"adr/backend/src/graphql/model"
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"adr/backend/src/service/userService"
	"adr/backend/src/utils"
	"context"
	"log"
)

func User(ctx context.Context, id string) (*model.User, error) {

	client := ctx.Value(srcModel.ConfigKey("client")).(*db.PrismaClient)

	data, err := userService.Find(id, client, ctx)
	if err != nil {
		log.Println("error userService Find", err)
		return nil, err
	}

	user := utils.UserConverter(data)

	return user, nil
}
