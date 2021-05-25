package infrastructure

import (
	"finder/graph"
	"finder/graph/generated"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func NewGraphQLHandler(r *graph.Resolver) *handler.Server {
	c := generated.Config{
		Resolvers: r,
	}
	// TODO: 今後実装する
	// c.Directives.Authentication = directives.Authentication

	return handler.NewDefaultServer(generated.NewExecutableSchema(c))
}

func NewPlayGroundHandler() http.Handler {
	return playground.Handler("GraphQL playground", "/query")
}
