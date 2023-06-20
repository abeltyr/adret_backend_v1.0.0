package companyService

import (
	"adr/backend/src/prisma/db"
	"context"
)

func Update(companyId string, ownerId string, client *db.PrismaClient, ctx context.Context) (*db.CompanyModel, error) {
	company, err := client.Company.FindUnique(db.Company.ID.Equals(companyId)).
		Update(
			db.Company.OwnerID.Set(ownerId),
		).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return company, nil

}
