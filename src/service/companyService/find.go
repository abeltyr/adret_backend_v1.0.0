package companyService

import (
	"adr/backend/src/prisma/db"
	"context"
)

func Find(id string, client *db.PrismaClient, ctx context.Context) (*db.CompanyModel, error) {

	company, err := client.Company.FindUnique(
		db.Company.ID.Equals(id),
	).Exec(ctx)

	return company, err
}
