package summaryRouter

import (
	"adr/backend/src/graphql/model"
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"adr/backend/src/service/orderService"
	"adr/backend/src/service/summaryService"
	"adr/backend/src/service/userService"
	"adr/backend/src/utils"
	"context"
	"errors"
	"log"
	"time"
)

func Summary(ctx context.Context, startDate string, endDate string) (*model.Summary, error) {

	client := ctx.Value(srcModel.ConfigKey("client")).(*db.PrismaClient)
	clientCtx := ctx.Value(srcModel.ConfigKey("clientCtx")).(context.Context)
	cognitoUser := ctx.Value(srcModel.ConfigKey("user"))

	start, _ := time.Parse("2006-01-02", startDate)
	end, _ := time.Parse("2006-01-02", endDate)
	now, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))

	if start.UnixMicro() > end.UnixMicro() {
		return nil, errors.New("start time can't be greater than the end")
	}

	// if start.UnixMicro() > now.UnixMicro() {
	// 	return nil, errors.New("can't get summary of the future")
	// }

	currentUser, currentCompany, fetchErr := userService.Checking(cognitoUser, srcModel.Manager, client, clientCtx)
	if fetchErr != nil {
		log.Println("error userService Checking", fetchErr)
		return nil, fetchErr
	}

	ctx = context.WithValue(ctx, srcModel.ConfigKey("currentUser"), currentUser)
	ctx = context.WithValue(ctx, srcModel.ConfigKey("currentCompany"), currentCompany)

	fetchSummary, _ := summaryService.Find(start, end, client, ctx)

	if start.UnixMicro() < now.UnixMicro() && (fetchSummary != nil && (fetchSummary.Earning > 0)) {
		summary := utils.SummaryConverter(fetchSummary)
		return summary, nil
	}

	earning, profit, err := orderService.Summary(start, end, nil, client, ctx)

	if err != nil {
		log.Println("error dailySummaryService Summary", err)
		return nil, err
	}

	summaryData, err := summaryService.Create(earning, profit, start, end, client, ctx)
	if err != nil {
		log.Println("error summaryService Create", err)
		return nil, err
	}

	summary := utils.SummaryConverter(summaryData)
	return summary, nil
}
