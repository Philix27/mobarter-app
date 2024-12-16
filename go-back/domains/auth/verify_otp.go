package auth

import (
	"errors"
	"mobarter/app"
	"mobarter/domains/crypto"
	"time"

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
				"token":   app.String,
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
							"email": &graphql.InputObjectFieldConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
						},
					},
				),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			args, err := app.ValidateArg("input", p)

			if err != nil {
				return nil, errors.New("Input required")
			}

			email := args["email"].(string)

			dto := VerifyOtpInput{
				Otp:   args["otp"].(string),
				Token: args["token"].(string),
			}

			if crypto.ValidateToken(dto.Token, dto.Otp) {
				return nil, errors.New("OTP is invalid")
			}

			accessToken, err := crypto.CreateJWTToken(email, time.Minute*30)
			if err != nil {
				return nil, errors.New("Could not get access token")
			}

			return map[string]interface{}{
				"message": "success",
				"token":   accessToken,
			}, nil
		},
	}
}
