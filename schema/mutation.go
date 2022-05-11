package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entify-schema-registry/config"
	"rxdrag.com/entify-schema-registry/consts"
)

var installInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "InstallInput",
		Fields: graphql.InputObjectConfigFieldMap{
			consts.DB_DRIVER: &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			consts.DB_HOST: &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			consts.DB_PORT: &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			consts.DB_SCHEMA: &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			consts.DB_USER: &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			consts.DB_PASSWORD: &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	},
)

func mutationFields() graphql.Fields {
	return graphql.Fields{
		"install": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"input": &graphql.ArgumentConfig{
					Type: &graphql.NonNull{
						OfType: installInputType,
					},
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return config.GetBool(consts.INSTALLED), nil
			},
		},
	}
}
