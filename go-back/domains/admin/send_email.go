package admin

import (
	"mobarter/app"

	"github.com/graphql-go/graphql"
)

func SendEmail(appState app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "Admin_SendEmailResponse",
			Fields: graphql.Fields{
				"message": app.String,
			},
		}),
		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type: graphql.NewInputObject(
					graphql.InputObjectConfig{
						Name: "Admin_SendEmail",
						Fields: graphql.InputObjectConfigFieldMap{
							"email":   app.InputString,
							"message": app.InputString,
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
