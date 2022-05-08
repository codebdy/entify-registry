package main

import "github.com/graphql-go/graphql"

var serviceType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Service",
		Fields: graphql.Fields{
			ID: &graphql.Field{
				Type: graphql.Int,
			},
			NAME: &graphql.Field{
				Type: graphql.String,
			},
			URL: &graphql.Field{
				Type: graphql.String,
			},
			TYPE_DEFS: &graphql.Field{
				Type:        graphql.String,
				Description: "Service types",
			},
		},
		Description: "Service type",
	},
)

func createSchema() (graphql.Schema, error) {
	fields := graphql.Fields{
		"services": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "Query", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	return graphql.NewSchema(schemaConfig)
}
