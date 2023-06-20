package companyRouter

import (
	"adr/backend/src/graphql/model"
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"adr/backend/src/service/companyBranchService"
	"adr/backend/src/service/companyService"
	"adr/backend/src/service/userService"
	"adr/backend/src/utils"
	"context"
	"log"
)

func Create(ctx context.Context, company model.CreateCompanyInput, owner *model.CreateOwnerInput, branch *model.CreateBranchInput) (*model.Company, error) {

	client := ctx.Value(srcModel.ConfigKey("client")).(*db.PrismaClient)

	createdCompany, err := companyService.Create(company, client, ctx)
	if err != nil {
		if createdCompany != nil {
			companyService.Delete(createdCompany.ID, client, ctx)
		}
		log.Println("error companyService create", err)
		return nil, err
	}

	branchData, err := companyBranchService.Create(companyBranchService.CompanyBranchInput{
		BranchName: branch.BranchName,
		Latitude:   branch.Latitude,
		Longitude:  branch.Longitude,
		CompanyId:  createdCompany.ID,
	}, client, ctx)
	if err != nil {
		if createdCompany != nil {
			companyService.Delete(createdCompany.ID, client, ctx)
		}
		log.Println("error companyService create", err)
		return nil, err
	}

	user, err := userService.Create(model.CreateUserInput{
		PhoneNumber: &owner.PhoneNumber,
		FullName:    &owner.FullName,
		UserName:    &owner.UserName,
		Password:    &owner.Password,
	}, srcModel.Manager, createdCompany, nil, client, ctx)
	if err != nil {
		if createdCompany != nil {
			companyService.Delete(createdCompany.ID, client, ctx)
		}
		if user != nil {
			userService.Delete(user.ID, client, ctx)
		}
		log.Println("error userService create", err)
		return nil, err
	}

	user, err = client.User.FindUnique(db.User.ID.Equals(user.ID)).Update(
		db.User.BranchID.Set(branchData.ID),
	).Exec(ctx)
	if err != nil {
		if createdCompany != nil {
			companyService.Delete(createdCompany.ID, client, ctx)
		}
		if user != nil {
			userService.Delete(user.ID, client, ctx)
		}
		log.Println("error companyService Update", err)
		return nil, err
	}

	companyUpdated, err := companyService.Update(createdCompany.ID, user.ID, client, ctx)
	if err != nil {
		if createdCompany != nil {
			companyService.Delete(createdCompany.ID, client, ctx)
		}
		if user != nil {
			userService.Delete(user.ID, client, ctx)
		}
		log.Println("error companyService Update", err)
		return nil, err
	}

	companyData := utils.CompanyConverter(companyUpdated)

	return companyData, nil
}
