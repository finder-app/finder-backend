package infrastructure

// import (
// 	"finder/graph"
// 	"finder/graph/directives"
// 	"finder/graph/generated"
// 	"net/http"

// 	"github.com/99designs/gqlgen/graphql/handler"
// 	"github.com/99designs/gqlgen/graphql/playground"
// )

// func NewGraphQLHandler(r *graph.Resolver) *handler.Server {
// 	c := generated.Config{
// 		Resolvers: r,
// 	}
// 	c.Directives.Authentication = directives.Authentication
// 	// 認証で弾かれると、↓は実行されない
// 	// fmt.Println(c.Directives.Authentication, "c.Directives.Authentication")
// 	return handler.NewDefaultServer(generated.NewExecutableSchema(c))
// }

// func NewPlayGroundHandler() http.Handler {
// 	return playground.Handler("GraphQL playground", "/query")
// }
