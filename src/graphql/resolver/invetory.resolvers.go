package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"adr/backend/src/graphql/generated"
	"adr/backend/src/graphql/model"
	"adr/backend/src/router/inventoryRouter"
	"adr/backend/src/router/productVariationRouter"
	"adr/backend/src/router/sharedRouter"
	"adr/backend/src/router/variationRouter"
	"context"
)

func (r *inventoryResolver) InventoryVariation(ctx context.Context, obj *model.Inventory) ([]*model.InventoryVariation, error) {
	return variationRouter.FindMany(ctx, obj.ID)
}

func (r *inventoryResolver) Media(ctx context.Context, obj *model.Inventory) (*string, error) {
	return inventoryRouter.Media(ctx, obj.ID)
}

func (r *inventoryResolver) Product(ctx context.Context, obj *model.Inventory) (*model.Product, error) {
	return sharedRouter.Product(ctx, *obj.ProductID)
}

func (r *inventoryVariationResolver) Title(ctx context.Context, obj *model.InventoryVariation) (*string, error) {
	if obj.ProductVariationID != nil {
		return productVariationRouter.FindTitle(ctx, *obj.ProductVariationID)
	}
	return nil, nil
}

func (r *queryResolver) Inventory(ctx context.Context, id string) (*model.Inventory, error) {
	return inventoryRouter.Find(ctx, id)
}

func (r *queryResolver) Inventories(ctx context.Context, input model.InventoriesFilter) ([]*model.Inventory, error) {
	return inventoryRouter.FindMany(ctx, input)
}

// Inventory returns generated.InventoryResolver implementation.
func (r *Resolver) Inventory() generated.InventoryResolver { return &inventoryResolver{r} }

// InventoryVariation returns generated.InventoryVariationResolver implementation.
func (r *Resolver) InventoryVariation() generated.InventoryVariationResolver {
	return &inventoryVariationResolver{r}
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type inventoryResolver struct{ *Resolver }
type inventoryVariationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) UpdateInventory(ctx context.Context, input model.UpdateInventoryInput) (bool, error) {
	return inventoryRouter.Update(ctx, input)
}
func (r *mutationResolver) DeleteInventory(ctx context.Context, id string) (*model.Inventory, error) {
	return nil, nil
}
