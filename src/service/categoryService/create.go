package categoryService

import (
	"adr/backend/src/prisma/db"
	"context"
	"strings"
)

func Create(name string, client *db.PrismaClient, ctx context.Context) (*db.CategoryModel, error) {
	category, err := client.Category.
		UpsertOne(
			db.Category.Name.Equals(strings.ToLower(name)),
		).
		Create(db.Category.Name.Set(strings.ToLower(name))).
		Update().
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	return category, nil

}
