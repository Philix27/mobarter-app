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
			"input": &graphql.ArgumentConfig{
				Type: graphql.NewInputObject(
					graphql.InputObjectConfig{
						Name: "Kyc_VerifyEmailInput",
						Fields: graphql.InputObjectConfigFieldMap{
							"email": &graphql.InputObjectFieldConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"otp": &graphql.InputObjectFieldConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"nin": &graphql.InputObjectFieldConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
						},
					},
				),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			return map[string]interface{}{
				"name": "Meat Pie",
			}, nil
		},
	}
}
