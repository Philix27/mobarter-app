package auth

import (
	"errors"
	"mobarter/app"
	"mobarter/domains/notification"

	"github.com/graphql-go/graphql"
)

type SendPhoneOtpDto struct {
	Phone   string
	Country string
}

func Auth_SendPhoneOtp(appState app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "Auth_SendPhoneOtpResponse",
			Fields: graphql.Fields{
				"message": app.String,
				"token":   app.String,
			},
		}),

		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type: graphql.NewInputObject(
					graphql.InputObjectConfig{
						Name: "SendPhoneOtpInput",
						Fields: graphql.InputObjectConfigFieldMap{
							"phone":   app.InputString,
							"country": app.InputString,
						},
					},
				),
			},
		},

		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			dto := SendPhoneOtpDto{
				Phone:   p.Args["phone"].(string),
				Country: p.Args["country"].(string),
			}

			err := notification.SendPhoneOtp(dto.Phone, dto.Country)

			if err != nil {
				return nil, errors.New("Could not send otp")
			}

			return map[string]interface{}{
				"message": "success",
			}, nil

		},
	}
}
