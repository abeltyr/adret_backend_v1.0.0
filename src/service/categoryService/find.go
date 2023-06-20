package categoryService

import (
	"adr/backend/src/prisma/db"
	"context"
	"strings"
)

func Find(name string, client *db.PrismaClient, ctx context.Context) (*db.CategoryModel, error) {

	category, err := client.Category.FindUnique(
		db.Category.Name.Equals(strings.ToLower(name)),
	).Exec(ctx)

	return category, err
}
