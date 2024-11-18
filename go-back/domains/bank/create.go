package bank

import (
	"mobarter/app"
	"mobarter/database"
	"mobarter/log"

	"github.com/graphql-go/graphql"
)

func CreateBankAccount(appState app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "CreateBankAccountResponse",
			Fields: graphql.Fields{
				"message": app.String,
			},
		}),
		Args: graphql.FieldConfigArgument{
			"accountName": app.ArgString,
			"accountNo":   app.ArgString,
			"bankName":    app.ArgString,
			"token":       app.ArgString,
			"wallet":      app.ArgString,
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			ilog := log.New("Create bank account")
			// todo - verify user credentials

			result := appState.DB.Create(&database.BankAccount{
				Name:     p.Args["accountName"].(string),
				No:       p.Args["accountNo"].(string),
				BankName: p.Args["bankName"].(string),
			})

			if result.Error != nil {
				println(result.Error)

				return map[string]interface{}{
					"error": result.Error,
				}, result.Error
			}
			ilog.Trace("Bank account added")
			return map[string]interface{}{
				"message": "success",
			}, nil
		},
	}
}
