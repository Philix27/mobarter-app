package auth

import (
	"errors"
	"mobarter/app"
	"mobarter/domains/crypto"

	"github.com/graphql-go/graphql"
)

type VerifyOtpInput struct {
	Token string
	Otp   string
}

func Auth_VerifyOtp(appState app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "Auth_VerifyOtpResponse",
			Fields: graphql.Fields{
				"message": app.String,
			},
		}),

		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type: graphql.NewInputObject(
					graphql.InputObjectConfig{
						Name: "Auth_VerifyOtpInput",
						Fields: graphql.InputObjectConfigFieldMap{
							"token": &graphql.InputObjectFieldConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"otp": &graphql.InputObjectFieldConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
						},
					},
				),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			dto := VerifyOtpInput{
				Otp:   p.Args["otp"].(string),
				Token: p.Args["token"].(string),
			}

			err := crypto.VerifyToken(dto.Token, dto.Otp)

			if err != nil {
				return nil, errors.New("OTP is invalid")
			}

			return map[string]interface{}{
				"message": "success",
			}, nil
		},
	}
}
