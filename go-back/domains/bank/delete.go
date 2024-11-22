package bank

import (
	"mobarter/app"

	"github.com/graphql-go/graphql"
)

func DeleteBankAccount(appState app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "DeleteBankAccountResponse",
			Fields: graphql.Fields{
				"message": app.String,
			},
		}),
		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type: graphql.NewInputObject(
					graphql.InputObjectConfig{
						Name: "DeleteBankAccountInput",
						Fields: graphql.InputObjectConfigFieldMap{
							"accountId": &graphql.InputObjectFieldConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"token": &graphql.InputObjectFieldConfig{
								Type: graphql.String,
							},
							"wallet": &graphql.InputObjectFieldConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
						},
					},
				),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// args, err := app.ValidateArg(optional.Some("input"), p)
			// if err != nil {
			// 	return nil, errors.New("Input required")
			// }
			// accountId := args["accountId"].(string)
			// token := args["token"].(string)

			return map[string]interface{}{
				"name": "Meat Pie",
			}, nil
		},
	}
}
