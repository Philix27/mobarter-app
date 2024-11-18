package user

import (
	"mobarter/app"
	"mobarter/database"

	"github.com/graphql-go/graphql"
)

func Create(appState app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "User_CreateResponse",
			Fields: graphql.Fields{
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
			// announcementItem := announcement{}

			res := appState.DB.Where(
				"walletAddress = ?", p.Args["walletAddress"].(string),
			).First(&database.User{})

			if res.Error != nil {
				appState.Logger.Error("CANNOT FIND_ONE:", res.Error)

				// todo: check if wallet already exist
				result := appState.DB.Create(&database.User{
					WalletAddress: p.Args["walletAddress"].(string),
					FirstName:     p.Args["firstName"].(string),
					LastName:      p.Args["lastName"].(string),
				})

				if result.Error != nil {
					println(result.Error)

					return map[string]interface{}{
						"error": result.Error,
					}, result.Error
				}

				return map[string]interface{}{
					"message": "success",
				}, nil

			} else {
				return map[string]interface{}{
					"error": "User already exist",
				}, res.Error
			}

		},
	}
}
