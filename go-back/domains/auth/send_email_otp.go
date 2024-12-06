package auth

import (
	"errors"
	"mobarter/app"
	"mobarter/domains/notification"

	"github.com/graphql-go/graphql"
)

type SendEmailOtpDto struct {
	Email string
}

func Auth_SendEmailOtp(appState app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "Auth_SendEmailOtpResponse",
			Fields: graphql.Fields{
				"message": app.String,
				"token":   app.String,
			},
		}),

		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type: graphql.NewInputObject(
					graphql.InputObjectConfig{
						Name: "SendEmailOtpInput",
						Fields: graphql.InputObjectConfigFieldMap{
							"email": app.InputStringOptional,
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

			dto := SendEmailOtpDto{
				Email: args["email"].(string),
			}

			println("After DT:", dto.Email)

			if notification.SendEmailOtp(dto.Email) != nil {
				return nil, errors.New("Could not send otp")
			}

			return map[string]interface{}{
				"message": "success",
			}, nil

		},
	}
}
