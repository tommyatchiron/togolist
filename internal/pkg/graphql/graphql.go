package graphql

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samsarahq/thunder/graphql"
	"github.com/samsarahq/thunder/graphql/graphiql"
	"github.com/samsarahq/thunder/graphql/introspection"
	"github.com/samsarahq/thunder/graphql/schemabuilder"
	"github.com/tommyatchiron/togolist/internal/pkg/router"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(newSchemaBuilder),
	fx.Provide(buildSchema),
	fx.Invoke(registerGraphQLRoutes),
)

func newSchemaBuilder() *schemabuilder.Schema {
	return schemabuilder.NewSchema()
}

func buildSchema(schema *schemabuilder.Schema) *graphql.Schema {
	s := schema.MustBuild()
	introspection.AddIntrospectionToSchema(s)
	return s
}

func registerGraphQLRoutes(s *graphql.Schema, r *router.Router) {
	r.RegisterApiRoutes("/graphql", func(rg *gin.RouterGroup) {
		rg.POST("", gin.WrapH(graphql.HTTPHandler(s)))
	})
	r.RegisterApiRoutes("/graphiql", func(rg *gin.RouterGroup) {
		rg.GET("*any", gin.WrapH(http.StripPrefix("/v1/graphiql/", graphiql.Handler())))
	})
}
