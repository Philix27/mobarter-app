package crm

import (
	"mobarter/app"

	"github.com/graphql-go/graphql"
)

func GetUserBankAccounts(appState app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "Admin_GetUserBankAccountsResponse",
			Fields: graphql.Fields{
				"message": app.String,
			},
		}),
		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type: graphql.NewInputObject(
					graphql.InputObjectConfig{
						Name: "Admin_GetUserBankAccountsInput",
						Fields: graphql.InputObjectConfigFieldMap{
							"userId": app.InputStringOptional,
						},
					},
				),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			//TODO -

			// args, err := app.ValidateArg(optional.Some("input"), p)
			// if err != nil {
			// 	return nil, errors.New("Input required")
			// }

			// //TODO - fetch user id
			// result := appState.DB.Create(&database.BankAccount{
			// 	Name:     args["accountName"].(string),
			// 	No:       args["accountNo"].(string),
			// 	BankName: args["bankName"].(string),
			// 	//TODO - use the user's id
			// 	UserID: 1,
			// })

			// if result.Error != nil {
			// 	println(result.Error)

			// 	return nil, errors.New("Could not create account")
			// }
			// ilog.Trace("Bank account added")
			return map[string]interface{}{
				"message": "success",
			}, nil
		},
	}
}
