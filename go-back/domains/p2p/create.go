package p2p

import (
	"errors"
	"mobarter/app"
	"mobarter/database"

	"github.com/graphql-go/graphql"
)

func CreateVendor(appState app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "CreateVendorResponse",
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
			// ilog := log.New("Create bank account")
			// todo - verify user credentials

			result := appState.DB.Create(&database.BankAccount{
				Name:     p.Args["accountName"].(string),
				No:       p.Args["accountNo"].(string),
				BankName: p.Args["bankName"].(string),
			})

			if result.Error != nil {
				println(result.Error)

				return map[string]interface{}{
					"error": "Could not create user",
				}, errors.New("Could not create user")
			}

			// ilog.Trace("Bank account added")
			return map[string]interface{}{
				"message":      "user created",
				"account_name": p.Args["accountName"].(string),
				"account_no":   p.Args["accountNo"].(string),
				"bank_name":    p.Args["bankName"].(string),
			}, nil
		},
	}
}
