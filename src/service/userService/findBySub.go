package userService

import (
	"adr/backend/src/prisma/db"
	"context"
)

func FindBySub(subsId string, client *db.PrismaClient, ctx context.Context) (*db.UserModel, error) {

	user, err := client.User.FindUnique(
		db.User.SubID.Equals(subsId),
	).Exec(ctx)

	return user, err
}
