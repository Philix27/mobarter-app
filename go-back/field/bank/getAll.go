package bank

import (
	"mobarter/app"

	"github.com/graphql-go/graphql"
)

func GetAllBankAccount(appState app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "getAllBankAccountResponse",
			Fields: graphql.Fields{
				"accountName": app.String,
				"accountNo":   app.String,
				"bankName":    app.String,
			},
		}),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			return map[string]interface{}{
				"name": "Meat Pie",
			}, nil
		},
	}
}
