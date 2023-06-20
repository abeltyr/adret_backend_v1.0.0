package categoryProductService

import (
	"adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"context"
	"strings"
)

func Find(filter model.CategoryProductInput, client *db.PrismaClient, ctx context.Context) (*db.CategoryProductModel, error) {

	parameter := []db.CategoryProductWhereParam{}
	if filter.CategoryName != "" {
		parameter = append(parameter,
			db.CategoryProduct.CategoryName.Equals(strings.ToLower(filter.CategoryName)))
	}

	if filter.ProductId != "" {
		parameter = append(parameter,
			db.CategoryProduct.ProductID.Equals(filter.ProductId))
	}

	FetchCategoryProduct, _ := client.CategoryProduct.
		FindFirst(
			parameter[:]...,
		).Exec(ctx)

	return FetchCategoryProduct, nil

}
