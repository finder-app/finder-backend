package directives

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/gqlerror"
)

func Authentication(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	err, ok := ctx.Value("AuthenticationError").(*gqlerror.Error)
	if ok {
		return nil, err
	}
	return next(ctx)
}
