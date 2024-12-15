package user

import (
	"errors"
	"mobarter/app"
	"mobarter/db"
	"mobarter/domains/crypto"

	"github.com/graphql-go/graphql"
)

type CreateDto struct {
	Email string
}

func Create(ap app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "User_CreateResponse",
			Fields: graphql.Fields{
				"message": app.String,
				"email":   app.String,
			},
		}),

		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type: graphql.NewInputObject(
					graphql.InputObjectConfig{
						Name: "CreateUserInput",
						Fields: graphql.InputObjectConfigFieldMap{
							"email": &graphql.InputObjectFieldConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"password": &graphql.InputObjectFieldConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"token": &graphql.InputObjectFieldConfig{
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
			token := args["token"].(string)

			err = crypto.ValidateToken(token, email)
			if err != nil {
				return nil, errors.New("Invalid token")
			}

			hash, err := crypto.HashPassword(password)
			if err != nil {
				return nil, errors.New("Could not hash password")
			}

			user, err := ap.DbQueries.User_Create(ap.Ctx, db.User_CreateParams{
				Email:          email,
				HashedPassword: hash,
			})

			if err != nil {
				println("Errox: " + err.Error())
				return nil, errors.New("User account could not be created")
			}

			return map[string]interface{}{
				"message": "success " + user.Email,
			}, nil

		},
	}
}
