package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"adr/backend/src/graphql/generated"
	"adr/backend/src/graphql/model"
	"adr/backend/src/router/sharedRouter"
	"adr/backend/src/router/userRouter"
	"context"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	return userRouter.Create(ctx, input)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUserInput) (*model.User, error) {
	return userRouter.Update(ctx, input)
}

func (r *mutationResolver) UpdateUserPassword(ctx context.Context, input model.UpdateUserPasswordInput) (*bool, error) {
	return userRouter.UpdatePassword(ctx, input)
}

func (r *mutationResolver) UpdatePersonalPassword(ctx context.Context, input model.UpdatePersonalPasswordInput) (*bool, error) {
	return userRouter.UpdatePersonalPassword(ctx, input)
}

func (r *queryResolver) CurrentUser(ctx context.Context) (*model.User, error) {
	return userRouter.FindMe(ctx)
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	return userRouter.Find(ctx, id)
}

func (r *queryResolver) Users(ctx context.Context, input model.UsersFilter) ([]*model.User, error) {
	return userRouter.FindMany(ctx, input)
}

func (r *userResolver) Company(ctx context.Context, obj *model.User) (*model.Company, error) {
	if obj.CompanyID != nil {
		return sharedRouter.Company(ctx, *obj.CompanyID)
	}
	return nil, nil
}

func (r *userResolver) Creator(ctx context.Context, obj *model.User) (*model.User, error) {
	if obj.CreatorID != nil {
		return sharedRouter.User(ctx, *obj.CreatorID)
	}
	return nil, nil
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
