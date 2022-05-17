package schema

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/mitchellh/mapstructure"
	"rxdrag.com/entify-schema-registry/config"
	"rxdrag.com/entify-schema-registry/consts"
	"rxdrag.com/entify-schema-registry/repository"
	"rxdrag.com/entify-schema-registry/utils"
)

const INPUT = "input"

var serviceInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "ServiceInput",
		Fields: graphql.InputObjectConfigFieldMap{
			consts.ID: &graphql.InputObjectFieldConfig{
				Type: &graphql.NonNull{
					OfType: graphql.Int,
				},
			},
			consts.URL: &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			consts.NAME: &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			consts.SERVICETYPE: &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	},
)

var installInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "InstallInput",
		Fields: graphql.InputObjectConfigFieldMap{
			consts.DB_DRIVER: &graphql.InputObjectFieldConfig{
				Type: &graphql.NonNull{
					OfType: graphql.String,
				},
			},
			consts.DB_HOST: &graphql.InputObjectFieldConfig{
				Type: &graphql.NonNull{
					OfType: graphql.String,
				},
			},
			consts.DB_PORT: &graphql.InputObjectFieldConfig{
				Type: &graphql.NonNull{
					OfType: graphql.String,
				},
			},
			consts.DB_DATABASE: &graphql.InputObjectFieldConfig{
				Type: &graphql.NonNull{
					OfType: graphql.String,
				},
			},
			consts.DB_USER: &graphql.InputObjectFieldConfig{
				Type: &graphql.NonNull{
					OfType: graphql.String,
				},
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
				INPUT: &graphql.ArgumentConfig{
					Type: &graphql.NonNull{
						OfType: installInputType,
					},
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				defer utils.PrintErrorStack()
				dbConfig := config.DbConfig{}
				mapstructure.Decode(p.Args[INPUT], &dbConfig)
				repository.Install(dbConfig)
				config.SetDbConfig(dbConfig)
				config.SetBool(consts.INSTALLED, true)
				config.WriteConfig()
				return config.GetBool(consts.INSTALLED), nil
			},
		},
		"addService": &graphql.Field{
			Type: serviceType,
			Args: graphql.FieldConfigArgument{
				INPUT: &graphql.ArgumentConfig{
					Type: &graphql.NonNull{
						OfType: serviceInputType,
					},
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				defer utils.PrintErrorStack()
				service := repository.Service{}
				mapstructure.Decode(p.Args[INPUT], &service)
				fmt.Println(service)
				repository.AddService(service)
				return repository.GetService(service.Id), nil
			},
		},
		"removeService": &graphql.Field{
			Type: serviceType,
			Args: graphql.FieldConfigArgument{
				consts.ID: &graphql.ArgumentConfig{
					Type: &graphql.NonNull{
						OfType: graphql.Int,
					},
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				defer utils.PrintErrorStack()
				service := repository.Service{}
				mapstructure.Decode(p.Args[INPUT], &service)
				oldService := repository.GetService(service.Id)
				repository.RemoveService(service.Id)
				return oldService, nil
			},
		},
		"updateService": &graphql.Field{
			Type: serviceType,
			Args: graphql.FieldConfigArgument{
				INPUT: &graphql.ArgumentConfig{
					Type: &graphql.NonNull{
						OfType: serviceInputType,
					},
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				defer utils.PrintErrorStack()
				service := repository.Service{}
				mapstructure.Decode(p.Args[INPUT], &service)
				repository.UpdateService(service)
				return repository.GetService(service.Id), nil
			},
		},
	}
}
