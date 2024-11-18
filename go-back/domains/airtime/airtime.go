package airtime

import (
	"mobarter/app"

	"github.com/graphql-go/graphql"
)

func PurchaseAirtime(appState app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "PurchaseAirtimeResponse",
			Fields: graphql.Fields{
				"message": app.String,
			},
		}),
		Args: graphql.FieldConfigArgument{
			"phone":           app.ArgString,
			"amount":          app.ArgString,
			"transactionHash": app.ArgInt,
			"network":         app.ArgString,
			"issuerAddress":   app.ArgString,
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			return map[string]interface{}{
				"name": "Meat Pie",
			}, nil
		},
	}
}
