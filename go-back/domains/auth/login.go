package auth

import (
	"errors"
	"mobarter/app"
	"mobarter/domains/crypto"

	"github.com/graphql-go/graphql"
)

type LoginInput struct {
	Email    string
	Password string
}

func Auth_Login(appState app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "Auth_LoginResponse",
			Fields: graphql.Fields{
				"message": app.String,
				"user_id": app.String,
				"token":   app.String,
			},
		}),

		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type: graphql.NewInputObject(
					graphql.InputObjectConfig{
						Name: "Auth_LoginInput",
						Fields: graphql.InputObjectConfigFieldMap{
							"email": &graphql.InputObjectFieldConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"password": &graphql.InputObjectFieldConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
						},
					},
				),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			dto := LoginInput{
				Email:    p.Args["email"].(string),
				Password: p.Args["password"].(string),
			}

			err := crypto.VerifyEmailPassword(dto.Email, dto.Password)

			if err != nil {
				return nil, errors.New("Invalid credentials")
			}

			return map[string]interface{}{
				"message": "success",
			}, nil
		},
	}
}
