package graph

import (
	"finder/graph/directives"
	"finder/graph/generated"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func NewGraphQLHandler(r *Resolver) *handler.Server {
	c := generated.Config{
		Resolvers: r,
	}
	// TODO: GraphQL側で認証。今後実装する
	c.Directives.Authentication = directives.Authentication

	return handler.NewDefaultServer(generated.NewExecutableSchema(c))
}

func NewPlayGroundHandler() http.Handler {
	return playground.Handler("GraphQL playground", "/query")
}
