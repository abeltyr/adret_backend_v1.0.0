package categoryService

import (
	"adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"context"
	"errors"
	"strings"
)

func Update(input model.CategoryUpdateInput, client *db.PrismaClient, ctx context.Context) (*db.CategoryModel, error) {

	category, err := Find(input.CurrentName, client, ctx)

	if err != nil {
		return nil, err
	}

	if category == nil {
		return nil, errors.New("invalid Name")
	}

	updated, err := client.Category.FindUnique(
		db.Category.Name.Equals(strings.ToLower(input.CurrentName)),
	).Update(
		db.Category.Name.Set(strings.ToLower(input.Name)),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return updated, err
}
