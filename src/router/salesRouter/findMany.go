package salesRouter

import (
	"adr/backend/src/graphql/model"
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"adr/backend/src/service/saleService"
	"adr/backend/src/service/userService"
	"context"
	"encoding/json"
	"log"
)

func FindMany(ctx context.Context, input model.SalesFilter) ([]*model.Sales, error) {

	client := ctx.Value(srcModel.ConfigKey("client")).(*db.PrismaClient)
	clientCtx := ctx.Value(srcModel.ConfigKey("clientCtx")).(context.Context)
	cognitoUser := ctx.Value(srcModel.ConfigKey("user"))

	_, currentCompany, fetchErr := userService.Checking(cognitoUser, "", client, clientCtx)
	if fetchErr != nil {
		log.Println("error userService Checking", fetchErr)
		return nil, fetchErr
	}

	ctx = context.WithValue(ctx, srcModel.ConfigKey("currentCompany"), currentCompany)

	salesData, err := saleService.FindMany(input, client, ctx)
	if err != nil {
		log.Println("error productService FindMany", err)
		return nil, err
	}

	j, _ := json.Marshal(salesData)

	var sales []*model.Sales
	_ = json.Unmarshal(j, &sales)

	return sales, nil
}
