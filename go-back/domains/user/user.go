package user

import (
	"mobarter/app"

	"github.com/graphql-go/graphql"
)

func UserInfo(appState app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "User_GetInfoResponse",
			Fields: graphql.Fields{
				"message":       app.String,
				"walletAddress": app.String,
				"firstName":     app.String,
				"lastName":      app.String,
				"id":      app.String,
			},
		}),
		Args: graphql.FieldConfigArgument{
			"walletAddress": app.ArgString,
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// ilog := log.New("Get user info")
			//
			res, err := findByWalletAddress(appState, p.Args["walletAddress"].(string))

			if err != nil {
				return map[string]interface{}{
					"message": "No user found",
				}, nil
			}

			return map[string]interface{}{
				"message":       "success",
				"id":            res.ID,
				"walletAddress": res.WalletAddress,
				"firstName":     res.FirstName,
				"lastName":      res.LastName,
			}, nil
		},
	}
}
