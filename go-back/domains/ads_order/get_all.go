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
				"id":            app.String,
			},
		}),

		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type: graphql.NewInputObject(
					graphql.InputObjectConfig{
						Name: "GetUserInput",
						Fields: graphql.InputObjectConfigFieldMap{
							"walletAddress": &graphql.InputObjectFieldConfig{
								Type: graphql.NewNonNull(graphql.Int),
							},
						},
					},
				),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			return map[string]interface{}{
				"message": "success",
			}, nil
		},
	}
}
