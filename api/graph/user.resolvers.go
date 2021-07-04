package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

// func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*domain.User, error) {
// 	panic(fmt.Errorf("not implemented"))
// }

// func (r *queryResolver) GetUsers(ctx context.Context) ([]*domain.User, error) {
// 	// NOTE: 仮のエラーハンドリング。// c.Directives.Authenticationを設定したら、それに移行する
// 	// if ctx.Value("currentUserUid") == nil {
// 	// 	return nil, &gqlerror.Error{
// 	// 		Message: "not token error",
// 	// 		Extensions: map[string]interface{}{
// 	// 			"status": http.StatusUnauthorized,
// 	// 		},
// 	// 	}
// 	// }
// 	// currentUserUid := ctx.Value("currentUserUid").(string)
// 	// return r.userUsecase.GetUsersByUid(currentUserUid)
// 	var users []*domain.User
// 	return users, nil
// }

// // Mutation returns generated.MutationResolver implementation.
// // func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// // // Query returns generated.QueryResolver implementation.
// // func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// // type mutationResolver struct{ *Resolver }
// // type queryResolver struct{ *Resolver }
