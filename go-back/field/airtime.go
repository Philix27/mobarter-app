package field

import "github.com/graphql-go/graphql"

var PurchaseAirtime *graphql.Field = &graphql.Field{
	Type: graphql.NewObject(graphql.ObjectConfig{
		Name: "PurchaseAirtimeResponse",
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
