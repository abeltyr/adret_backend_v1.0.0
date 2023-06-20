package summaryService

import (
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"context"
	"time"
)

func Find(startDate time.Time, endDate time.Time, client *db.PrismaClient, ctx context.Context) (*db.SummaryModel, error) {

	company := ctx.Value(srcModel.ConfigKey("currentCompany")).(*db.CompanyModel)

	summary, err := client.Summary.
		FindUnique(
			db.Summary.CompanyIDStartDateEndDate(
				db.Summary.CompanyID.Equals(company.ID),
				db.Summary.StartDate.Equals(startDate),
				db.Summary.EndDate.Equals(endDate),
			),
		).
		Exec(ctx)

	return summary, err
}
