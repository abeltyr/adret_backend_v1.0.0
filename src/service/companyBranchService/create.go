package companyBranchService

import (
	"adr/backend/src/prisma/db"
	"context"
)

func Create(input CompanyBranchInput, client *db.PrismaClient, ctx context.Context) (*db.CompanyBranchModel, error) {

	companyBranch, err := client.CompanyBranch.UpsertOne(
		db.CompanyBranch.BranchNameCompanyID(
			db.CompanyBranch.BranchName.Equals(input.BranchName),
			db.CompanyBranch.CompanyID.Equals(input.CompanyId),
		),
	).Create(
		db.CompanyBranch.BranchName.Set(input.BranchName),
		db.CompanyBranch.Company.Link(db.Company.ID.Equals(input.CompanyId)),
		db.CompanyBranch.Latitude.Set(input.Latitude),
		db.CompanyBranch.Longitude.Set(input.Longitude),
	).Update(
		db.CompanyBranch.BranchName.Set(input.BranchName),
		db.CompanyBranch.Company.Link(db.Company.ID.Equals(input.CompanyId)),
		db.CompanyBranch.Latitude.Set(input.Latitude),
		db.CompanyBranch.Longitude.Set(input.Longitude)).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return companyBranch, nil

}
