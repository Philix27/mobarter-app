package user

import (
	"errors"
	"mobarter/app"

	"github.com/graphql-go/graphql"
)

type CreateDto struct {
	Email string
}

func Create(appState app.AppState) *graphql.Field {
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

			// dto := CreateDto{
			// 	Email: p.Args["email"].(string),
			// }

			user, err := appState.DbQueries.User_Create(appState.Ctx, email)

			if err != nil {

				println("Oops an error occurred " + email)
				println("Errox: " + err.Error())
			}
			println("Hello user " + email)
			// db.CreateUser()
			// todo: check if wallet already exist

			// err = createUserRepo(appState, &dto)

			// if err != nil {
			// 	return nil, errors.New("Could not create user")
			// }

			return map[string]interface{}{
				"message": "success " + user.Email,
			}, nil

		},
	}
}
