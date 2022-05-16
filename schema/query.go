package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entify-schema-registry/config"
	"rxdrag.com/entify-schema-registry/consts"
	"rxdrag.com/entify-schema-registry/repository"
	"rxdrag.com/entify-schema-registry/utils"
)

func queryFields() graphql.Fields {
	return graphql.Fields{
		"services": &graphql.Field{
			Type: &graphql.NonNull{
				OfType: &graphql.List{
					OfType: serviceType,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				defer utils.PrintErrorStack()
				return repository.GetServices(), nil
			},
		},
		"installed": &graphql.Field{
			Type: graphql.Boolean,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				defer utils.PrintErrorStack()
				return config.GetBool(consts.INSTALLED), nil
			},
		},
		"authenticationService": &graphql.Field{
			Type: serviceType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				defer utils.PrintErrorStack()
				service := repository.GetAuthService()
				if service != nil {
					return *service, nil
				}
				return nil, nil
			},
		},
	}
}
