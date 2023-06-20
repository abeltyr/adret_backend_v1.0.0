package categoryCompanyService

import (
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"context"
	"strings"
)

func Create(categoryName string, client *db.PrismaClient, ctx context.Context) (*db.CategoryCompanyModel, error) {

	company := ctx.Value(srcModel.ConfigKey("currentCompany")).(*db.CompanyModel)
	category := strings.ToLower(categoryName)
	createdCategoryCompany, err := client.CategoryCompany.
		UpsertOne(
			db.CategoryCompany.CategoryNameCompanyID(
				db.CategoryCompany.CategoryName.Equals(category),
				db.CategoryCompany.CompanyID.Equals(company.ID),
			),
		).
		Create(
			db.CategoryCompany.Category.Link(db.Category.Name.Equals(category)),
			db.CategoryCompany.Company.Link(db.Company.ID.Equals(company.ID)),
		).
		Update().
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	return createdCategoryCompany, nil

}
