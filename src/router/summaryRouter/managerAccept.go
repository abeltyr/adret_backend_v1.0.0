package summaryRouter

import (
	"adr/backend/src/graphql/model"
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"adr/backend/src/service/employeeDailySummaryService"
	"adr/backend/src/service/userService"
	"adr/backend/src/utils"
	"context"
	"log"
)

func ManagerAccept(ctx context.Context, id string) (*model.Summary, error) {
	client := ctx.Value(srcModel.ConfigKey("client")).(*db.PrismaClient)
	clientCtx := ctx.Value(srcModel.ConfigKey("clientCtx")).(context.Context)
	cognitoUser := ctx.Value(srcModel.ConfigKey("user"))

	currentUser, _, fetchErr := userService.Checking(cognitoUser, "", client, clientCtx)
	if fetchErr != nil {
		log.Println("error userService Checking", fetchErr)
		return nil, fetchErr
	}

	ctx = context.WithValue(ctx, srcModel.ConfigKey("currentUser"), currentUser)

	summaryData, err := employeeDailySummaryService.Update(id, client, ctx)
	if err != nil {
		log.Println("error employeeDailySummeryService Update", err)
		return nil, err
	}

	summary := utils.SummaryConverter(summaryData)
	return summary, nil
}
