package field

import (
	"mobarter/app"

	"github.com/graphql-go/graphql"
)

func UserInfo(appState app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "UserResponse",
			Fields: graphql.Fields{
				"walletAddress": String,
				"firstName":     String,
				"lastName":      String,
				"country":       String,
			},
		}),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			return map[string]interface{}{
				"name": "Meat Pie",
			}, nil
		},
	}
}
