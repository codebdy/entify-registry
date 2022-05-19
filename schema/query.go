package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entify-schema-registry/config"
	"rxdrag.com/entify-schema-registry/utils"
)

func queryFields() graphql.Fields {
	return graphql.Fields{
		"authUrl": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				defer utils.PrintErrorStack()
				return config.AuthUrl(), nil
			},
		},
	}
}
