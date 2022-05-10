package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entify-schema-registry/config"
	"rxdrag.com/entify-schema-registry/consts"
	"rxdrag.com/entify-schema-registry/repository"
)

var serviceType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Service",
		Fields: graphql.Fields{
			consts.ID: &graphql.Field{
				Type: graphql.Int,
			},
			consts.NAME: &graphql.Field{
				Type: graphql.String,
			},
			consts.URL: &graphql.Field{
				Type: graphql.String,
			},
			consts.SERVICETYPE: &graphql.Field{
				Type: graphql.String,
			},
			consts.TYPE_DEFS: &graphql.Field{
				Type: graphql.String,
			},
			consts.IS_ALIVE: &graphql.Field{
				Type: graphql.Boolean,
			},
			consts.VERSION: &graphql.Field{
				Type: graphql.String,
			},
			consts.ADDED_TIME: &graphql.Field{
				Type: graphql.DateTime,
			},
			consts.UPDATED_TIME: &graphql.Field{
				Type: graphql.DateTime,
			},
		},
		Description: "Service type",
	},
)

func CreateSchema() (graphql.Schema, error) {
	fields := graphql.Fields{
		"services": &graphql.Field{
			Type: &graphql.NonNull{
				OfType: &graphql.List{
					OfType: serviceType,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return repository.GetServices(), nil
			},
		},
		"installed": &graphql.Field{
			Type: graphql.Boolean,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return config.GetBool(consts.IS_INSTALLED), nil
			},
		},
		"authenticationService": &graphql.Field{
			Type: serviceType,
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "Query", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	return graphql.NewSchema(schemaConfig)
}
