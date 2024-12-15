package auth

import (
	"errors"
	"mobarter/app"
	"mobarter/domains/crypto"
	"strconv"
	"time"

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
			args, err := app.ValidateArg("input", p)

			if err != nil {
				return nil, errors.New("Input required")
			}

			email := args["email"].(string)
			password := args["password"].(string)

			user, err := appState.DbQueries.User_GetByEmail(appState.Ctx, email)

			if err != nil {
				return nil, errors.New("Invalid email and password")
			}

			if crypto.ComparePassword(password, user.HashedPassword) {
				return nil, errors.New("Invalid credentials")
			}

			token, err := crypto.CreateJWTToken(strconv.Itoa(int(user.ID)), time.Hour*24*7)

			if err != nil {
				return nil, errors.New("Could not generate token")
			}

			return map[string]interface{}{
				"message": "success",
				"token":   token,
			}, nil
		},
	}
}
