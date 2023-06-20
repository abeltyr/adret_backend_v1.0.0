package companyService

import (
	"adr/backend/src/graphql/model"
	"adr/backend/src/prisma/db"
	"context"
	"strings"
)

func Create(input model.CreateCompanyInput, client *db.PrismaClient, ctx context.Context) (*db.CompanyModel, error) {

	companyCode := strings.ToLower(input.CompanyCode)

	company, err := client.Company.CreateOne(
		db.Company.Name.Set(input.Name),
		db.Company.CompanyCode.Set(companyCode),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return company, nil

}
