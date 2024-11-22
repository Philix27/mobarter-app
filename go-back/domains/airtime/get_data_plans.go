package airtime

import (
	"errors"
	"mobarter/app"

	"github.com/LNMMusic/optional"
	"github.com/graphql-go/graphql"
)

type GetDataPlansDto struct {
	category string
	network  string
}

func GetDataPlans(appState app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "GetDataPlansResponse",
			Fields: graphql.Fields{
				"message": app.String,
			},
		}),
		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type: graphql.NewInputObject(
					graphql.InputObjectConfig{
						Name: "GetDataPlansInput",
						Fields: graphql.InputObjectConfigFieldMap{

							"network": &graphql.InputObjectFieldConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"category": &graphql.InputObjectFieldConfig{
								Type: graphql.String,
							},
						},
					},
				),
			},
		},

		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// ilog := log.New("Purchase airtime")
			args, err := app.ValidateArg(optional.Some("input"), p)
			if err != nil {
				return nil, errors.New("Input required")
			}
			dto := &GetDataPlansDto{
				network:  args["network"].(string),
				category: args["category"].(string),
			}

			if !verifyNetwork(dto.network) {
				return nil, errors.New("Invalid network provider")
			}

			// todo: log
			// ilog.Debug("Recharged airtime for" + dto.phone)

			return map[string]interface{}{
				"message": "success",
			}, nil
		},
	}
}
