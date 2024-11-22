package user

import (
	"errors"
	"mobarter/app"

	"github.com/graphql-go/graphql"
)

type CreateDto struct {
	WalletAddress string
	FirstName     string
	LastName      string
}

func Create(appState app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "User_CreateResponse",
			Fields: graphql.Fields{
				"message":       app.String,
				"walletAddress": app.String,
				"firstName":     app.String,
				"lastName":      app.String,
				"id":            app.String,
			},
		}),

		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type: graphql.NewInputObject(
					graphql.InputObjectConfig{
						Name: "CreateUserInput",
						Fields: graphql.InputObjectConfigFieldMap{
							"walletAddress": &graphql.InputObjectFieldConfig{
								Type: graphql.NewNonNull(graphql.Int),
							},
							"firstName": &graphql.InputObjectFieldConfig{
								Type: graphql.String,
							},
							"lastName": &graphql.InputObjectFieldConfig{
								Type: graphql.String,
							},
						},
					},
				),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			dto := CreateDto{
				WalletAddress: p.Args["walletAddress"].(string),
				FirstName:     p.Args["firstName"].(string),
				LastName:      p.Args["lastName"].(string),
			}

			// todo: check if wallet already exist

			_, err := findByWalletAddress(appState, dto.WalletAddress)

			if err == nil {

				return nil, errors.New("User already exist")
			}

			err = createUserRepo(appState, &dto)

			if err != nil {
				return nil, errors.New("Could not create user")
			}

			return map[string]interface{}{
				"message":       "success",
				"walletAddress": p.Args["walletAddress"].(string),
				"firstName":     p.Args["firstName"].(string),
				"lastName":      p.Args["lastName"].(string),
			}, nil

		},
	}
}
