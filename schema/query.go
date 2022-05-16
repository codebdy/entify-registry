package schema

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entify-schema-registry/config"
	"rxdrag.com/entify-schema-registry/consts"
	"rxdrag.com/entify-schema-registry/repository"
	"rxdrag.com/entify-schema-registry/utils"
)

func covertService(service *repository.Service) map[string]interface{} {
	return map[string]interface{}{
		"id":          service.Id,
		"name":        service.Name,
		"url":         service.Url,
		"serviceType": service.ServiceType.String,
		"typeDefs":    service.TypeDefs.String,
		"isAlive":     service.IsAlive.Bool,
		"version":     service.Version.String,
		"addedTime":   service.AddedTime.Time,
		"updatedTime": service.UpdatedTime.Time,
	}
}

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
				services := repository.GetServices()
				datas := []map[string]interface{}{}
				for i := range services {
					datas = append(datas, covertService(services[i]))
				}
				return datas, nil
			},
		},
		"status": &graphql.Field{
			Type: graphql.NewObject(
				graphql.ObjectConfig{
					Name: "ServiceStatus",
					Fields: graphql.Fields{
						consts.INSTALLED: &graphql.Field{
							Type: graphql.Boolean,
						},
						consts.AUTH_INSTALLED: &graphql.Field{
							Type: graphql.Boolean,
						},
					},
					Description: "Service status",
				},
			),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				defer utils.PrintErrorStack()
				return map[string]interface{}{
					consts.INSTALLED:      config.GetBool(consts.INSTALLED),
					consts.AUTH_INSTALLED: repository.GetAuthService() != nil,
				}, nil
			},
		},
		"authenticationService": &graphql.Field{
			Type: serviceType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				defer utils.PrintErrorStack()
				service := repository.GetAuthService()
				if service != nil {
					return covertService(service), nil
				}
				return nil, nil
			},
		},
	}
}
