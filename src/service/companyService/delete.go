package companyService

import (
	"adr/backend/src/prisma/db"
	"context"
)

func Delete(id string, client *db.PrismaClient, ctx context.Context) (*db.CompanyModel, error) {
	company, err := client.Company.FindUnique(
		db.Company.ID.Equals(id),
	).Delete().Exec(ctx)

	if err != nil {
		return nil, err
	}

	return company, nil

}
