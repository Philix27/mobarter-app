package user

import (
	"mobarter/app"
	"mobarter/log"

	"github.com/graphql-go/graphql"
)

func UserInfo(appState app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "User_GetInfoResponse",
			Fields: graphql.Fields{
				"walletAddress": app.String,
				"firstName":     app.String,
				"lastName":      app.String,
			},
		}),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			ilog := log.New("Get user info")
			// 
			ilog.Trace("Deleted bank account")
			return map[string]interface{}{
				"name": "Meat Pie",
			}, nil
		},
	}
}
