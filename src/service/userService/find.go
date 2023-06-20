package userService

import (
	"adr/backend/src/prisma/db"
	"context"
)

func Find(id string, client *db.PrismaClient, ctx context.Context) (*db.UserModel, error) {

	user, err := client.User.FindUnique(
		db.User.ID.Equals(id),
	).Exec(ctx)

	return user, err
}
