package userService

import (
	"adr/backend/src/prisma/db"
	"context"
)

func FindByUsername(username string, client *db.PrismaClient, ctx context.Context) (*db.UserModel, error) {

	user, err := client.User.FindUnique(
		db.User.UserName.Equals(username),
	).Exec(ctx)

	return user, err
}
