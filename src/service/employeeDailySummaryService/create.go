package employeeDailySummaryService

import (
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"context"
	"time"
)

func Create(earning float64, profit float64, date time.Time, employeeID string, client *db.PrismaClient, ctx context.Context) (*db.EmployeeDailySummaryModel, error) {

	company := ctx.Value(srcModel.ConfigKey("currentCompany")).(*db.CompanyModel)

	summary, err := client.EmployeeDailySummary.
		UpsertOne(
			db.EmployeeDailySummary.CompanyIDDateEmployeeID(
				db.EmployeeDailySummary.CompanyID.Equals(company.ID),
				db.EmployeeDailySummary.Date.Equals(date),
				db.EmployeeDailySummary.EmployeeID.Equals(employeeID),
			),
		).
		Create(
			db.EmployeeDailySummary.Earning.Set(earning),
			db.EmployeeDailySummary.Profit.Set(profit),
			db.EmployeeDailySummary.Date.Set(date),
			db.EmployeeDailySummary.EmployeeID.Set(employeeID),
			db.EmployeeDailySummary.Company.Link(db.Company.ID.Equals(company.ID)),
		).
		Update(
			db.EmployeeDailySummary.Earning.Set(earning),
			db.EmployeeDailySummary.Profit.Set(profit)).
		Exec(ctx)

	return summary, err

}
