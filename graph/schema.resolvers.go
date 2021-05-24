package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"finder/graph/generated"
	"finder/graph/model"
	"fmt"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Text:   input.Text,
		ID:     "1",
		UserID: input.UserID,
		// User: &model.User{ID: input.UserID, Name: "user " + input.UserID},
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := &model.User{
		ID:   "3",
		Name: input.Name,
	}
	if err := r.DB.Table("user").Create(user).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}
	return user, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	// NOTE: テーブル変えたらバグってるから使うな
	var todos []*model.Todo
	// if err := r.DB.Table("todos").Preload("User2").Find(&todos).Error; err != nil {
	// 	fmt.Println(err)
	// 	return nil, err
	// }
	return todos, nil
}

func (r *queryResolver) Todo(ctx context.Context, id string) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	users := []*model.User{}
	if err := r.DB.Table("users2").Preload("Todos").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	user := &model.User{}
	if err := r.DB.Table("users2").Where("id = ?", id).Preload("Todos").Take(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
