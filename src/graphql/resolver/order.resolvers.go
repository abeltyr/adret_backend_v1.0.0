package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"adr/backend/src/graphql/generated"
	"adr/backend/src/graphql/model"
	"adr/backend/src/router/orderRouter"
	"adr/backend/src/router/sharedRouter"
	"context"
)

func (r *mutationResolver) CreateLocalOrder(ctx context.Context, input model.CreateLocalOrderInput) (*model.Order, error) {
	return orderRouter.Create(ctx, input)
}

func (r *orderResolver) Seller(ctx context.Context, obj *model.Order) (*model.User, error) {
	if obj.SellerID != nil {
		return sharedRouter.User(ctx, *obj.SellerID)
	}
	return nil, nil
}

func (r *orderResolver) Company(ctx context.Context, obj *model.Order) (*model.Company, error) {
	if obj.CompanyID != nil {
		return sharedRouter.Company(ctx, *obj.CompanyID)
	}
	return nil, nil
}

func (r *orderResolver) OnlineOrderDetail(ctx context.Context, obj *model.Order) (*model.OnlineOrderDetail, error) {
	// if obj.SellerID != nil {
	// 	return sharedRouter.User(ctx, *obj.SellerID)
	// }
	return nil, nil
}

func (r *orderResolver) OnlineOrderPayment(ctx context.Context, obj *model.Order) (*model.OnlineOrderPayment, error) {
	// if obj.SellerID != nil {
	// 	return sharedRouter.User(ctx, *obj.SellerID)
	// }
	return nil, nil
}

func (r *orderResolver) Sales(ctx context.Context, obj *model.Order) ([]*model.Sales, error) {
	return orderRouter.Sales(ctx, obj.ID)
}

func (r *queryResolver) Order(ctx context.Context, id string) (*model.Order, error) {
	return orderRouter.Find(ctx, id)
}

func (r *queryResolver) Orders(ctx context.Context, input model.OrdersFilter) ([]*model.Order, error) {
	return orderRouter.FindMany(ctx, input)
}

// Order returns generated.OrderResolver implementation.
func (r *Resolver) Order() generated.OrderResolver { return &orderResolver{r} }

type orderResolver struct{ *Resolver }
