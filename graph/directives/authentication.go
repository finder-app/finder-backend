package directives

import (
	"context"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/gqlerror"
)

func Authentication(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	err, ok := ctx.Value("AuthenticationError").(*gqlerror.Error)
	// fmt.Println(err.Extensions, "err.Extensions") // map[status:401] err.Extensions
	if ok {
		// returnする時にerr.statusは含まれてる
		// return nil, err

		// ここに書いてもreturnされない！！！
		hoge := &gqlerror.Error{
			Message: err.Error(),
			Extensions: map[string]interface{}{
				"status": http.StatusUnauthorized,
			},
		}
		return nil, hoge
	}
	return next(ctx)
}
