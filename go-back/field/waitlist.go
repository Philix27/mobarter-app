package field

import "github.com/graphql-go/graphql"

var joinWaitList *graphql.Field = &graphql.Field{
	Type: graphql.NewObject(graphql.ObjectConfig{
		Name: "waitListResponse",
		Fields: graphql.Fields{
			"message": String,
		},
	}),
	Args: graphql.FieldConfigArgument{
		"email":     argString,
		"firstName": argString,
		"lastName":  argString,
		"country":   argString,
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {

		return map[string]interface{}{
			"name": "Meat Pie",
		}, nil
	},
}
