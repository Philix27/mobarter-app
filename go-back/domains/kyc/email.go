package kyc

import (
	"mobarter/app"

	"github.com/graphql-go/graphql"
)

func Kyc_VerifyEmail(appState app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "Kyc_VerifyEmailResponse",
			Fields: graphql.Fields{
				"message": app.String,
			},
		}),
		Args: graphql.FieldConfigArgument{
			"email": app.ArgString,
			"otp":   app.ArgString,
			"token": app.ArgString,
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			return map[string]interface{}{
				"name": "Meat Pie",
			}, nil
		},
	}
}
