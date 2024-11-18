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
			"accountName": app.ArgString,
			"accountNo":   app.ArgString,
			"bankName":    app.ArgString,
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			return map[string]interface{}{
				"name": "Meat Pie",
			}, nil
		},
	}
}
