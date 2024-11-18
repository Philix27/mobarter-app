package field

import (
	"mobarter/app"

	"github.com/graphql-go/graphql"
)

func Auth_SentOtp(appState app.AppState) *graphql.Field {
	return &graphql.Field{
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
			// appState.DB
			return map[string]interface{}{
				"name": "Meat Pie",
			}, nil
		},
	}
}
