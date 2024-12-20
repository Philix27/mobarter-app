package field

import (
	"fmt"
	"mobarter/app"

	"github.com/graphql-go/graphql"
)

func CountryList(appState app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "User_CreateResponse",
			Fields: graphql.Fields{
				"name": app.String,
				"flag": app.String,
			},
		}),
		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type: graphql.NewInputObject(
					graphql.InputObjectConfig{
						Name: "CountryInput",
					},
				),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			fmt.Println("Hit country list")
			return map[string]interface{}{
				"name": "Meat Pie",
			}, nil
		},
	}
}
