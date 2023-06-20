package categoryService

import (
	"adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"adr/backend/src/utils"
	"context"
)

func FindMany(filter model.FilterData, client *db.PrismaClient, ctx context.Context) ([]db.CategoryModel, error) {

	limit := utils.LimitSetter(filter.Limit)

	parameter := []db.CategoryWhereParam{}

	FetchCategory := client.Category.
		FindMany(
			parameter[:]...,
		).
		Take(int(limit)).OrderBy(
		db.Category.CreatedAt.Order(db.DESC),
	)

	if filter.After != "" {
		FetchCategory = client.Category.
			FindMany(
				parameter[:]...,
			).
			Take(int(limit)).
			Skip(1).
			Cursor(db.Category.Name.Cursor(string(filter.After))).
			OrderBy(
				db.Category.CreatedAt.Order(db.DESC),
			)
	}

	if filter.Before != "" {
		FetchCategory = client.Category.
			FindMany(
				parameter[:]...,
			).
			Take(int(limit)).
			Skip(1).
			Cursor(db.Category.Name.Cursor(string(filter.Before))).
			OrderBy(
				db.Category.CreatedAt.Order(db.ASC),
			)
	}

	users, err := FetchCategory.
		Exec(ctx)

	return users, err

}
