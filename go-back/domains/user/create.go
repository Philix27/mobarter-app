package user

import (
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
								Type: graphql.NewNonNull(graphql.Int),
							},
						},
					},
				),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			dto := CreateDto{
				Email: p.Args["email"].(string),
			}

			// todo: check if wallet already exist

			// err = createUserRepo(appState, &dto)

			// if err != nil {
			// 	return nil, errors.New("Could not create user")
			// }

			return map[string]interface{}{
				"message": "success",
				"email":   dto.Email,
			}, nil

		},
	}
}
