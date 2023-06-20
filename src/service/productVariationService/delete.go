package productVariationService

import (
	"adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"context"
)

func Delete(input model.CategoryCompanyInput, client *db.PrismaClient, ctx context.Context) (*db.CategoryCompanyModel, error) {

	deleted, err := client.CategoryCompany.FindUnique(
		db.CategoryCompany.CategoryNameCompanyID(
			db.CategoryCompany.CategoryName.Equals(input.CategoryName),
			db.CategoryCompany.CompanyID.Equals(input.CompanyId),
		),
	).Delete().Exec(ctx)

	if err != nil {
		return nil, err
	}

	return deleted, err
}
