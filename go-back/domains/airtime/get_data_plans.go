package airtime

import (
	"mobarter/app"

	"github.com/graphql-go/graphql"
)

var dataPlans = []map[string]interface{}{
	{"id": "1", "name": "John Doe"},
	{"id": "2", "name": "Jane Smith"},
	{"id": "3", "name": "Alice Johnson"},
}

type GetDataPlansDto struct {
	category string
	network  string
}

func GetDataPlans(appState app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "GetDataPlansResponse",
			Fields: graphql.Fields{
				// "message": app.String,
				"id":   &graphql.Field{Type: graphql.String},
				"name": &graphql.Field{Type: graphql.String},
				"dataPlans": &graphql.Field{
					Type: graphql.String,
					Args: graphql.FieldConfigArgument{
						"name": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
				},
			},
		}),
		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type: graphql.NewInputObject(
					graphql.InputObjectConfig{
						Name: "GetDataPlansInput",
						Fields: graphql.InputObjectConfigFieldMap{
							"network":  networkProvider,
							"category": app.InputStringOptional,
						},
					},
				),
			},
		},

		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// ilog := log.New("Purchase airtime")
			// args, err := app.ValidateArg(optional.Some("input"), p)
			// if err != nil {
			// 	return nil, errors.New("Input required")
			// }
			// dto := &GetDataPlansDto{
			// 	network:  args["network"].(string),
			// 	category: args["category"].(string),
			// }

			// if !verifyNetwork(dto.network) {
			// 	return nil, errors.New("Invalid network provider")
			// }

			// todo: log
			// ilog.Debug("Recharged airtime for" + dto.phone)

			return dataPlans, nil
		},
	}
}
