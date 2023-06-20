package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"adr/backend/src/graphql/generated"
	"adr/backend/src/graphql/model"
	"adr/backend/src/router/sharedRouter"
	"context"
)

func (r *companyResolver) Owner(ctx context.Context, obj *model.Company) (*model.User, error) {
	if obj.OwnerID != nil {
		return sharedRouter.User(ctx, *obj.OwnerID)
	}
	return nil, nil
}

func (r *mutationResolver) CreateOwnerCompany(ctx context.Context, company model.CreateCompanyInput, owner *model.CreateOwnerInput, branch *model.CreateBranchInput) (*model.Company, error) {
	// return companyRouter.Create(ctx, company, owner, branch)
	return nil, nil
}

func (r *mutationResolver) ResetOwnerPassword(ctx context.Context, username string, newPassword string) (bool, error) {
	// return userRouter.UpdateAdminPassword(ctx, username, newPassword)
	return false, nil
}

// Company returns generated.CompanyResolver implementation.
func (r *Resolver) Company() generated.CompanyResolver { return &companyResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type companyResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
