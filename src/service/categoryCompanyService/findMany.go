package categoryCompanyService

import (
	"adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"adr/backend/src/utils"
	"context"
	"strings"
)

func FindMany(filter model.CategoryCompanyFindManyFilter, client *db.PrismaClient, ctx context.Context) ([]db.CategoryCompanyModel, error) {

	limit := utils.LimitSetter(filter.Filter.Limit)

	parameter := []db.CategoryCompanyWhereParam{}
	if filter.CategoryCompany.CategoryName != "" {
		parameter = append(parameter,
			db.CategoryCompany.CategoryName.Equals(strings.ToLower(filter.CategoryCompany.CategoryName)))
	}

	if filter.CategoryCompany.CompanyId != "" {
		parameter = append(parameter,
			db.CategoryCompany.CompanyID.Equals(filter.CategoryCompany.CompanyId))
	}

	FetchCategoryCompany := client.CategoryCompany.
		FindMany(
			parameter[:]...,
		).
		Take(int(limit)).
		OrderBy(
			db.CategoryCompany.CategoryName.Order(db.DESC),
		)

	if filter.Filter.After != "" {
		FetchCategoryCompany = client.CategoryCompany.
			FindMany(
				parameter[:]...,
			).
			Take(int(limit)).
			Skip(1).
			Cursor(db.CategoryCompany.CategoryName.Cursor(string(filter.Filter.After))).
			OrderBy(
				db.CategoryCompany.CategoryName.Order(db.DESC),
			)
	}

	if filter.Filter.Before != "" {
		FetchCategoryCompany = client.CategoryCompany.
			FindMany(
				parameter[:]...,
			).
			Take(int(limit)).
			Skip(1).
			Cursor(db.CategoryCompany.CategoryName.Cursor(string(filter.Filter.Before))).
			OrderBy(
				db.CategoryCompany.CategoryName.Order(db.DESC),
			)
	}

	users, err := FetchCategoryCompany.
		Exec(ctx)

	return users, err

}
