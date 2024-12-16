package user

import (
	"errors"
	"mobarter/app"
	"mobarter/db"
	"mobarter/domains/crypto"

	"github.com/graphql-go/graphql"
)

func ResetPassword(ap app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "User_ResetPassword",
			Fields: graphql.Fields{
				"message": app.String,
			},
		}),

		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type: graphql.NewInputObject(
					graphql.InputObjectConfig{
						Name: "ResetPasswordInput",
						Fields: graphql.InputObjectConfigFieldMap{
							"email":    app.InputString,
							"password": app.InputString,
							"token":    app.InputString,
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
			// extract response data
			email := args["email"].(string)
			password := args["password"].(string)
			token := args["token"].(string)

			if crypto.ValidateToken(token, email) {
				return nil, errors.New("Invalid token")
			}

			hash, err := crypto.HashPassword(password)
			if err != nil {
				return nil, errors.New("Could not hash password")
			}

			user, err := ap.DbQueries.User_GetByEmail(ap.Ctx, email)
			if err != nil {
				return nil, errors.New("User not found")
			}

			err = ap.DbQueries.User_ResetPassword(ap.Ctx, db.User_ResetPasswordParams{
				ID:             user.ID,
				HashedPassword: hash,
			})

			if err != nil {
				println("Errox: " + err.Error())
				return nil, errors.New("Could not reset password")
			}

			return map[string]interface{}{
				"message": "success " + user.Email,
			}, nil

		},
	}
}
