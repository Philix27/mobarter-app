package field

import "github.com/graphql-go/graphql"

var UserInfo *graphql.Field = &graphql.Field{
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
