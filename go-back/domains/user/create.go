package user

import (
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
			"walletAddress": app.ArgString,
			"firstName":     app.ArgString,
			"lastName":      app.ArgString,
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			dto := CreateDto{
				WalletAddress: p.Args["walletAddress"].(string),
				FirstName:     p.Args["firstName"].(string),
				LastName:      p.Args["lastName"].(string),
			}

			// todo: check if wallet already exist

			_, err := findByWalletAddress(appState, dto.WalletAddress)
			// res := appState.DB.Where(
			// 	"wallet_address = ?", dto.WalletAddress,
			// ).First(&database.User{})

			if err == nil {

				return map[string]interface{}{
					"message": "User already exist",
				}, err
			}

			err = createUserRepo(appState, &dto)

			if err != nil {
				return map[string]interface{}{
					"message": "could not create user",
				}, err
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
