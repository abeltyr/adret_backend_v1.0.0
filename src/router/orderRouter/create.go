package orderRouter

import (
	"adr/backend/src/graphql/model"
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"adr/backend/src/service/inventoryService"
	"adr/backend/src/service/orderService"
	"adr/backend/src/service/saleService"
	"adr/backend/src/service/summaryService"
	"adr/backend/src/service/userService"
	"adr/backend/src/utils"
	"context"
	"errors"
	"log"
	"time"
)

func Create(ctx context.Context, input model.CreateLocalOrderInput) (*model.Order, error) {

	client := ctx.Value(srcModel.ConfigKey("client")).(*db.PrismaClient)
	clientCtx := ctx.Value(srcModel.ConfigKey("clientCtx")).(context.Context)
	cognitoUser := ctx.Value(srcModel.ConfigKey("user"))

	currentUser, currentCompany, fetchErr := userService.Checking(cognitoUser, "", client, clientCtx)
	if fetchErr != nil {
		log.Println("error userService Checking", fetchErr)
		return nil, fetchErr
	}
	ctx = context.WithValue(ctx, srcModel.ConfigKey("currentUser"), currentUser)
	ctx = context.WithValue(ctx, srcModel.ConfigKey("currentCompany"), currentCompany)

	now, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))

	summary, _ := summaryService.FindByDate(now, client, ctx)
	if currentUser.UserRole != db.RoleManager && (summary != nil && (summary.ManagerAccepted)) {
		log.Println("today sales collected and closed by manger")
		return nil, errors.New("today sales collected and closed by manger")
	}

	orderData, err := orderService.Create(input.Note, true, client, ctx)
	if err != nil {
		log.Println("error orderService Create", err)
		return nil, err
	}

	ctx = context.WithValue(ctx, srcModel.ConfigKey("order"), orderData)

	totalPrice := 0.0
	totalProfit := 0.0
	for _, sale := range input.SalesInput {

		inventory, err := inventoryService.Find(sale.InventoryID, client, ctx)

		if err != nil {
			orderService.Delete(orderData.ID, client, ctx)
			log.Println("error orderService Delete", err)
			return nil, err
		}
		errorData := errors.New("inventory sell amount can't be zero or less")
		if sale.Amount <= 0 {
			orderService.Delete(orderData.ID, client, ctx)
			log.Println("amount less than zero", errorData)
			return nil, errorData
		}
		if sale.SellingPrice <= 0 {
			errorData = errors.New("inventory sellingPrice price can't be zero or less")
			orderService.Delete(orderData.ID, client, ctx)
			log.Println("sellingPrice less than zero", errorData)
			return nil, errorData
		}
		if sale.SellingPrice < inventory.MinSellingPriceEstimation {
			errorData = errors.New("can't sell under the min selling price")
			log.Println(errorData)
			orderService.Delete(orderData.ID, client, ctx)
			return nil, errorData
		}
		if inventory.Available < inventory.SalesAmount+sale.Amount {
			errorData = errors.New("can't sell over the available amount")
			log.Println("error out of stock", errorData)
			orderService.Delete(orderData.ID, client, ctx)
			return nil, errorData
		}
		saleData, err := saleService.Create(*sale, inventory, client, ctx)
		if err != nil {
			orderService.Delete(orderData.ID, client, ctx)
			log.Println("error saleService Create", err)
			return nil, err
		}
		_, err = inventoryService.SaleAmountIncrement(sale.InventoryID, sale.Amount, client, ctx)
		if err != nil {
			orderService.Delete(orderData.ID, client, ctx)
			log.Println("error inventoryService SaleAmountIncrement", err)
			return nil, err
		}
		totalPrice = totalPrice + (saleData.SellingPrice * float64(sale.Amount))
		totalProfit = totalProfit + (saleData.Profit * float64(sale.Amount))
	}

	orderData, err = orderService.Update(orderData.ID, totalPrice, totalProfit, client, ctx)
	if err != nil {
		log.Println("error orderService Create", err)
		return nil, err
	}

	order := utils.OrderConverter(orderData)

	return order, nil
}
