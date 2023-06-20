package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"adr/backend/src/graphql/generated"
	"adr/backend/src/graphql/model"
	"adr/backend/src/router/salesRouter"
	"adr/backend/src/router/sharedRouter"
	"context"
	"fmt"
)

func (r *queryResolver) Sale(ctx context.Context, orderID string, inventoryID string) (*model.Sales, error) {
	return salesRouter.Find(ctx, orderID, inventoryID)
}

func (r *queryResolver) Sales(ctx context.Context, input model.SalesFilter) ([]*model.Sales, error) {
	return salesRouter.FindMany(ctx, input)
}

func (r *salesResolver) Inventory(ctx context.Context, obj *model.Sales) (*model.Inventory, error) {
	if obj.InventoryID != nil {
		return sharedRouter.Inventory(ctx, *obj.InventoryID)
	}
	return nil, nil
}

func (r *salesResolver) Order(ctx context.Context, obj *model.Sales) (*model.Order, error) {
	panic(fmt.Errorf("not implemented"))
}

// Sales returns generated.SalesResolver implementation.
func (r *Resolver) Sales() generated.SalesResolver { return &salesResolver{r} }

type salesResolver struct{ *Resolver }
