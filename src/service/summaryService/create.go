package summaryService

import (
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"context"
	"time"
)

func Create(earning float64, profit float64, startDate time.Time, endDate time.Time, client *db.PrismaClient, ctx context.Context) (*db.SummaryModel, error) {

	company := ctx.Value(srcModel.ConfigKey("currentCompany")).(*db.CompanyModel)

	summary, err := client.Summary.
		UpsertOne(
			db.Summary.CompanyIDStartDateEndDate(
				db.Summary.CompanyID.Equals(company.ID),
				db.Summary.StartDate.Equals(startDate),
				db.Summary.EndDate.Equals(endDate),
			),
		).
		Create(
			db.Summary.Earning.Set(earning),
			db.Summary.Profit.Set(profit),
			db.Summary.StartDate.Set(startDate),
			db.Summary.EndDate.Set(endDate),
			db.Summary.Company.Link(db.Company.ID.Equals(company.ID)),
		).
		Update(
			db.Summary.Earning.Set(earning),
			db.Summary.Profit.Set(profit)).
		Exec(ctx)

	return summary, err

}
