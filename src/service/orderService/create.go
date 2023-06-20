package orderService

import (
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"adr/backend/src/service/tableCountService"
	"context"
	"fmt"
	"time"
)

func Create(noteData *string, paid bool, client *db.PrismaClient, ctx context.Context) (*db.OrderModel, error) {

	currentCompany := ctx.Value(srcModel.ConfigKey("currentCompany")).(*db.CompanyModel)
	currentUser := ctx.Value(srcModel.ConfigKey("currentUser")).(*db.UserModel)

	add := 1
	count, err := tableCountService.Create("orderCount", &add, client, ctx)
	if err != nil {
		return nil, err
	}

	note := ""
	if noteData != nil {
		note = *noteData
	}

	orderCount := fmt.Sprintf(`%v-%d`, currentCompany.CompanyCode, (1000 + int64(count.Count)))

	date := time.Now().Format("2006-01-02")
	currentDate, _ := time.Parse("2006-01-02", date)

	createdOrder, err := client.Order.CreateOne(
		db.Order.Online.Set(false),
		db.Order.OrderNumber.Set(orderCount),
		db.Order.Note.Set(note),
		db.Order.TotalPrice.Set(0),
		db.Order.TotalProfit.Set(0),
		db.Order.Date.Set(currentDate),
		db.Order.SellerID.Set(currentUser.ID),
		db.Order.CompanyID.Set(currentCompany.ID),
		db.Order.Paid.Set(paid),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return createdOrder, nil

}
