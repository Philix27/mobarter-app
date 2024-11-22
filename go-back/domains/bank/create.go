package bank

import (
	"errors"
	"mobarter/app"
	"mobarter/database"

	"github.com/LNMMusic/optional"
	"github.com/graphql-go/graphql"
)

var users = []map[string]interface{}{
	{"id": "1", "name": "John Doe", "email": "john@example.com", "age": 30},
	{"id": "2", "name": "Jane Smith", "email": "jane@example.com", "age": 25},
}

func CreateBankAccount(appState app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "CreateBankAccountResponse",
			Fields: graphql.Fields{
				"message": app.String,
			},
		}),
		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type: graphql.NewInputObject(
					graphql.InputObjectConfig{
						Name: "CreateBankAccountInput",
						Fields: graphql.InputObjectConfigFieldMap{
							"accountName": &graphql.InputObjectFieldConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"accountNo": &graphql.InputObjectFieldConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"bankName": &graphql.InputObjectFieldConfig{
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
			//TODO -

			args, err := app.ValidateArg(optional.Some("input"), p)
			if err != nil {
				return nil, errors.New("Input required")
			}

			//TODO - fetch user id
			result := appState.DB.Create(&database.BankAccount{
				Name:     args["accountName"].(string),
				No:       args["accountNo"].(string),
				BankName: args["bankName"].(string),
				//TODO - use the user's id
				UserID: 1,
			})

			if result.Error != nil {
				println(result.Error)

				return nil, errors.New("Could not create account")
			}
			// ilog.Trace("Bank account added")
			return map[string]interface{}{
				"message": "success",
			}, nil
		},
	}
}
