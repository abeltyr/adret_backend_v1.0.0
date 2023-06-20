package summaryService

import (
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"context"
	"time"
)

func FindByDate(date time.Time, client *db.PrismaClient, ctx context.Context) (*db.EmployeeDailySummaryModel, error) {

	company := ctx.Value(srcModel.ConfigKey("currentCompany")).(*db.CompanyModel)
	currentUser := ctx.Value(srcModel.ConfigKey("currentUser")).(*db.UserModel)

	sales, err := client.EmployeeDailySummary.FindUnique(
		db.EmployeeDailySummary.CompanyIDDateEmployeeID(
			db.EmployeeDailySummary.CompanyID.Equals(company.ID),
			db.EmployeeDailySummary.Date.Equals(date),
			db.EmployeeDailySummary.EmployeeID.Equals(currentUser.ID),
		),
	).Exec(ctx)

	return sales, err
}
