package field

import "github.com/graphql-go/graphql"

var Auth_SentOtp *graphql.Field = &graphql.Field{
	Type: graphql.NewObject(graphql.ObjectConfig{
		Name: "Auth_SentOtpResponse",
		Fields: graphql.Fields{
			"message": String,
		},
	}),
	Args: graphql.FieldConfigArgument{
		"email":           argString,
		"phone":           argString,
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
