package bank

import (
	"mobarter/app"

	"github.com/graphql-go/graphql"
)

func GetBankAccount(appState app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "GetBankAccountResponse",
			Fields: graphql.Fields{
				"name":           app.String,
				"id":             app.Int,
				"dexId":          app.Int,
				"upVotes":        app.Int,
				"downVotes":      app.Int,
				"totalVotes":     app.Int,
				"winPercentage":  app.Float,
				"lossPercentage": app.Float,
			},
		}),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			return map[string]interface{}{
				"name": "Meat Pie",
			}, nil
		},
	}
}
