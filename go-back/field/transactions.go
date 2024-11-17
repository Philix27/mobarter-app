package field

import (
	"mobarter/app"

	"github.com/graphql-go/graphql"
)

var getAllTransactionsResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "getAllTransactionsResponse",
	Fields: graphql.Fields{
		"walletAddress": String,
		"firstName":     String,
		"lastName":      String,
		"country":       String,
	},
})

func GetAllTransactions(appState app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: getAllTransactionsResponse,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			return map[string]interface{}{
				"name": "Meat Pie",
			}, nil
		},
	}
}

func GetOneTransactions(appState app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "getOneTransactionsResponse",
			Fields: graphql.Fields{
				"walletAddress": String,
				"firstName":     String,
				"lastName":      String,
				"country":       String,
			},
		}),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			return map[string]interface{}{
				"name": "Meat Pie",
			}, nil
		},
	}
}
