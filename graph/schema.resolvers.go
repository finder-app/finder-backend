package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"finder/graph/generated"
	"finder/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	// panic(fmt.Errorf("not implemented"))
	return &model.Todo{
		ID:   "todo001",
		Text: "部屋の掃除",
		Done: false,
		User: &model.User{
			ID:   "user001",
			Name: "たろー",
		},
	}, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	// panic(fmt.Errorf("not implemented"))
	return []*model.Todo{
		{
			ID:   "todo001",
			Text: "部屋の掃除",
			Done: false,
			User: &model.User{
				ID:   "user001",
				Name: "たろー",
			},
		},
		{
			ID:   "todo002",
			Text: "買い物",
			Done: true,
			User: &model.User{
				ID:   "user001",
				Name: "たろー",
			},
		},
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
