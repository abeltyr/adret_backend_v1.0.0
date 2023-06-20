package summaryRouter

import (
	"adr/backend/src/graphql/model"
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"adr/backend/src/service/employeeDailySummaryService"
	"adr/backend/src/service/orderService"
	"adr/backend/src/service/userService"
	"adr/backend/src/utils"
	"context"
	"errors"
	"log"
	"time"
)

func EmployeeDailySummary(ctx context.Context, input model.EmployeeDailySummaryFilter) (*model.Summary, error) {
	client := ctx.Value(srcModel.ConfigKey("client")).(*db.PrismaClient)
	clientCtx := ctx.Value(srcModel.ConfigKey("clientCtx")).(context.Context)
	cognitoUser := ctx.Value(srcModel.ConfigKey("user"))

	start, _ := time.Parse("2006-01-02", input.Date)
	now, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))

	// if start.UnixMicro() > now.UnixMicro() {
	// 	return nil, errors.New("can't get summary of the future")
	// }

	currentUser, currentCompany, fetchErr := userService.Checking(cognitoUser, "", client, clientCtx)
	if fetchErr != nil {
		log.Println("error userService Checking", fetchErr)
		return nil, fetchErr
	}

	if currentUser != nil && (currentUser.UserRole != db.RoleManager && input.EmployeeID != currentUser.ID) {
		return nil, errors.New("only admin and specific employee can fetch the following data")
	}

	ctx = context.WithValue(ctx, srcModel.ConfigKey("currentUser"), currentUser)
	ctx = context.WithValue(ctx, srcModel.ConfigKey("currentCompany"), currentCompany)

	fetchSummary, _ := employeeDailySummaryService.Find(start, input.EmployeeID, client, ctx)

	if start.UnixMicro() < now.UnixMicro() && (fetchSummary != nil && (fetchSummary.Earning > 0)) {
		log.Println("using fetchSummary")
		summary := utils.SummaryConverter(fetchSummary)
		return summary, nil
	}

	if fetchSummary != nil && fetchSummary.ManagerAccepted {
		summary := utils.SummaryConverter(fetchSummary)
		return summary, nil
	}

	earning, profit, err := orderService.Summary(start, start, &input.EmployeeID, client, ctx)

	if err != nil {
		log.Println("error summaryService Summary", err)
		return nil, err
	}

	summaryData, err := employeeDailySummaryService.Create(earning, profit, start, input.EmployeeID, client, ctx)
	if err != nil {
		log.Println("error summaryService Create", err)
		return nil, err
	}

	summary := utils.SummaryConverter(summaryData)
	return summary, nil
}
