package kyc

import (
	"mobarter/app"

	"github.com/graphql-go/graphql"
)

func Kyc_GetLevel(appState app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "Kyc_GetLevelResponse",
			Fields: graphql.Fields{
				"accountName": app.String,
				"accountNo":   app.String,
				"bankName":    app.String,
			},
		}),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			return map[string]interface{}{
				"name": "Meat Pie",
			}, nil
		},
	}
}
