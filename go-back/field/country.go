package field

import "github.com/graphql-go/graphql"

var CountryList *graphql.Field = &graphql.Field{
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
