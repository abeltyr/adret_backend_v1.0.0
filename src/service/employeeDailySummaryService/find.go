package employeeDailySummaryService

import (
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"context"
	"time"
)

func Find(date time.Time, employeeID string, client *db.PrismaClient, ctx context.Context) (*db.EmployeeDailySummaryModel, error) {

	company := ctx.Value(srcModel.ConfigKey("currentCompany")).(*db.CompanyModel)

	summary, err := client.EmployeeDailySummary.
		FindUnique(
			db.EmployeeDailySummary.CompanyIDDateEmployeeID(
				db.EmployeeDailySummary.CompanyID.Equals(company.ID),
				db.EmployeeDailySummary.Date.Equals(date),
				db.EmployeeDailySummary.EmployeeID.Equals(employeeID),
			),
		).
		Exec(ctx)

	return summary, err
}
