package employeeDailySummaryService

import (
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"context"
)

func Update(id string, client *db.PrismaClient, ctx context.Context) (*db.EmployeeDailySummaryModel, error) {

	user := ctx.Value(srcModel.ConfigKey("currentUser")).(*db.UserModel)

	print(id)
	summary, err := client.EmployeeDailySummary.
		FindUnique(
			db.EmployeeDailySummary.ID.Equals(id),
		).
		Update(
			db.EmployeeDailySummary.ManagerAccepted.Set(true),
			db.EmployeeDailySummary.ManagerID.Set(user.ID),
		).
		Exec(ctx)

	return summary, err

}
