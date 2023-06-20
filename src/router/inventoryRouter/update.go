package inventoryRouter

import (
	"adr/backend/src/graphql/model"
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"adr/backend/src/service/inventoryService"
	"adr/backend/src/service/productService"
	"adr/backend/src/service/userService"
	"context"
	"errors"
	"log"
)

func Update(ctx context.Context, input model.UpdateInventoryInput) (bool, error) {

	client := ctx.Value(srcModel.ConfigKey("client")).(*db.PrismaClient)
	clientCtx := ctx.Value(srcModel.ConfigKey("clientCtx")).(context.Context)
	cognitoUser := ctx.Value(srcModel.ConfigKey("user"))

	_, company, fetchErr := userService.Checking(cognitoUser, "", client, clientCtx)
	if fetchErr != nil {
		log.Println("error userService Checking", fetchErr)
		return false, fetchErr
	}

	product, err := productService.Find(input.ProductID, client, ctx)
	if err != nil {
		log.Println("error productService Find", err)
		return false, err
	}

	companyId, ok := product.CompanyID()
	if !ok {
		log.Println("error productService Find", err)
		return false, errors.New("product doesn't have a company")
	}

	if companyId != company.ID {
		log.Println("error productService Find", err)
		return false, errors.New("you can't update another company product")
	}

	inventoryData, err := inventoryService.UpdateMany(input, client, ctx)
	if err != nil {
		log.Println("error inventoryService Update", err)
		return false, err
	}

	return inventoryData, nil
}
