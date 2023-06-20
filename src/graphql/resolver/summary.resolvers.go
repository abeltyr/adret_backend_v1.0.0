package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"adr/backend/src/graphql/generated"
	"adr/backend/src/graphql/model"
	"adr/backend/src/router/sharedRouter"
	"adr/backend/src/router/summaryRouter"
	"context"
)

func (r *mutationResolver) ManagerAccept(ctx context.Context, id string) (*model.Summary, error) {
	return summaryRouter.ManagerAccept(ctx, id)
}

func (r *queryResolver) Summary(ctx context.Context, startDate string, endDate string) (*model.Summary, error) {
	return summaryRouter.Summary(ctx, startDate, endDate)
}

func (r *queryResolver) EmployeeDailySummary(ctx context.Context, input model.EmployeeDailySummaryFilter) (*model.Summary, error) {
	return summaryRouter.EmployeeDailySummary(ctx, input)
}

func (r *summaryResolver) Manager(ctx context.Context, obj *model.Summary) (*model.User, error) {
	if obj.ManagerID != nil {
		return sharedRouter.User(ctx, *obj.ManagerID)
	}
	return nil, nil
}

func (r *summaryResolver) Employee(ctx context.Context, obj *model.Summary) (*model.User, error) {
	if obj.EmployeeID != nil {
		return sharedRouter.User(ctx, *obj.EmployeeID)
	}
	return nil, nil
}

func (r *summaryResolver) SummaryInventory(ctx context.Context, obj *model.Summary, filter *model.FilterInput) ([]*model.Inventory, error) {
	return summaryRouter.InventoryList(ctx, obj, filter)
}

// Summary returns generated.SummaryResolver implementation.
func (r *Resolver) Summary() generated.SummaryResolver { return &summaryResolver{r} }

type summaryResolver struct{ *Resolver }
