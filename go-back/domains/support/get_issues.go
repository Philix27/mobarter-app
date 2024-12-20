package support

import (
	"mobarter/app"

	"github.com/graphql-go/graphql"
)

func GetIssueTopics(appState app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "Support_GetIssueTopicsResponse",
			Fields: graphql.Fields{
				"message": app.String,
				"issues": &graphql.Field{
					Type: graphql.NewList(
						graphql.NewObject(
							graphql.ObjectConfig{
								Name: "Issues",
								Fields: graphql.Fields{
									"id":    app.String,
									"title": app.String,
								},
							},
						),
					),
				},
			},
		}),
		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type: graphql.NewInputObject(
					graphql.InputObjectConfig{
						Name: "Support_GetIssueTopicsInput",
						Fields: graphql.InputObjectConfigFieldMap{
							"token": app.InputString,
						},
					},
				),
			},
		},

		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			return map[string]interface{}{
				"message":      "user created",
				"account_name": p.Args["accountName"].(string),
				"account_no":   p.Args["accountNo"].(string),
				"bank_name":    p.Args["bankName"].(string),
			}, nil
		},
	}
}
