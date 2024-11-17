package field

import (
	"mobarter/app"

	"github.com/graphql-go/graphql"
)

func CountryList(appState app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "CountryResponse",
			Fields: graphql.Fields{
				"name": String,
				"flag": String,
			},
		}),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			return map[string]interface{}{
				"name": "Meat Pie",
			}, nil
		},
	}
}
