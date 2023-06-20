package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"adr/backend/src/graphql/generated"
	"adr/backend/src/graphql/model"
	"adr/backend/src/router/productRouter"
	"adr/backend/src/router/sharedRouter"
	"context"
)

func (r *mutationResolver) CreateProduct(ctx context.Context, input model.CreateProductInput) (*model.Product, error) {
	return productRouter.Create(ctx, input)
}

func (r *mutationResolver) UpdateProduct(ctx context.Context, input model.UpdateProductInput) (*model.Product, error) {
	return productRouter.Update(ctx, input)
}

func (r *productResolver) Creator(ctx context.Context, obj *model.Product) (*model.User, error) {
	if obj.CreatorID != nil {
		return sharedRouter.User(ctx, *obj.CreatorID)
	}
	return nil, nil
}

func (r *productResolver) Company(ctx context.Context, obj *model.Product) (*model.Company, error) {
	if obj.CompanyID != nil {
		return sharedRouter.Company(ctx, *obj.CompanyID)
	}
	return nil, nil
}

func (r *productResolver) Category(ctx context.Context, obj *model.Product) (*string, error) {
	return productRouter.Category(ctx, obj)
}

func (r *productResolver) InStock(ctx context.Context, obj *model.Product) (*int, error) {
	return productRouter.Remaining(ctx, obj.ID)
}

func (r *productResolver) Media(ctx context.Context, obj *model.Product) ([]*string, error) {
	return productRouter.Media(ctx, obj)
}

func (r *productResolver) Inventory(ctx context.Context, obj *model.Product) ([]*model.Inventory, error) {
	return productRouter.Inventory(ctx, obj)
}

func (r *productResolver) Variation(ctx context.Context, obj *model.Product) ([]*model.ProductVariation, error) {
	// panic(fmt.Errorf("not implemented"))
	// if obj.CreatorID != nil {
	// 	return sharedRouter.User(ctx, *obj.CreatorID)
	// }
	return nil, nil
}

func (r *queryResolver) Product(ctx context.Context, id string) (*model.Product, error) {
	return productRouter.Find(ctx, id)
}

func (r *queryResolver) ProductByCode(ctx context.Context, productCode string) (*model.Product, error) {
	return productRouter.FindByCode(ctx, productCode)
}

func (r *queryResolver) Products(ctx context.Context, input model.ProductsFilter) ([]*model.Product, error) {
	return productRouter.FindMany(ctx, input)
}

// Product returns generated.ProductResolver implementation.
func (r *Resolver) Product() generated.ProductResolver { return &productResolver{r} }

type productResolver struct{ *Resolver }
