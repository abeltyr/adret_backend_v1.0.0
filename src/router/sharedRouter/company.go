package sharedRouter

import (
	"adr/backend/src/graphql/model"
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"adr/backend/src/service/companyService"
	"adr/backend/src/utils"
	"context"
	"log"
)

func Company(ctx context.Context, id string) (*model.Company, error) {

	client := ctx.Value(srcModel.ConfigKey("client")).(*db.PrismaClient)

	data, err := companyService.Find(id, client, ctx)
	if err != nil {
		log.Println("error companyService Find", err)
		return nil, err
	}

	company := utils.CompanyConverter(data)

	return company, nil
}
