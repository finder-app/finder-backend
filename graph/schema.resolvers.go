package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"finder/domain"
	"finder/graph/generated"
	"finder/graph/model"
	"fmt"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*domain.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetUsers(ctx context.Context) ([]*domain.User, error) {
	currentUserUid := ctx.Value("currentUserUid").(string)
	return r.userUsecase.GetUsersByUid(currentUserUid)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
