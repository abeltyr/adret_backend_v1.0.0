package categoryProductService

import (
	"adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"adr/backend/src/utils"
	"context"
	"strings"
)

func FindMany(filter model.CategoryProductFindManyFilter, client *db.PrismaClient, ctx context.Context) ([]db.CategoryProductModel, error) {

	limit := utils.LimitSetter(filter.Filter.Limit)

	parameter := []db.CategoryProductWhereParam{}
	if filter.CategoryProduct.CategoryName != "" {
		parameter = append(parameter,
			db.CategoryProduct.CategoryName.Equals(strings.ToLower(filter.CategoryProduct.CategoryName)))
	}

	if filter.CategoryProduct.ProductId != "" {
		parameter = append(parameter,
			db.CategoryProduct.ProductID.Equals(filter.CategoryProduct.ProductId))
	}

	FetchCategoryProduct := client.CategoryProduct.
		FindMany(
			parameter[:]...,
		).
		Take(int(limit)).
		OrderBy(
			db.CategoryProduct.CategoryName.Order(db.DESC),
		)

	if filter.Filter.After != "" {
		FetchCategoryProduct = client.CategoryProduct.
			FindMany(
				parameter[:]...,
			).
			Take(int(limit)).
			Skip(1).
			Cursor(db.CategoryProduct.CategoryName.Cursor(string(filter.Filter.After))).
			OrderBy(
				db.CategoryProduct.CategoryName.Order(db.DESC),
			)
	}

	if filter.Filter.Before != "" {
		FetchCategoryProduct = client.CategoryProduct.
			FindMany(
				parameter[:]...,
			).
			Take(int(limit)).
			Skip(1).
			Cursor(db.CategoryProduct.CategoryName.Cursor(string(filter.Filter.Before))).
			OrderBy(
				db.CategoryProduct.CategoryName.Order(db.DESC),
			)
	}

	users, err := FetchCategoryProduct.
		Exec(ctx)

	return users, err

}
