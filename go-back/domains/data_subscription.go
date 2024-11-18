package field

import (
	"mobarter/app"

	"github.com/graphql-go/graphql"
)

func PurchaseData(appState app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "PurchaseDataResponse",
			Fields: graphql.Fields{
				"message": String,
			},
		}),
		Args: graphql.FieldConfigArgument{
			"phone":           argString,
			"amount":          argString,
			"transactionHash": argInt,
			"network":         argString,
			"issuerAddress":   argString,
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			return map[string]interface{}{
				"name": "Meat Pie",
			}, nil
		},
	}
}