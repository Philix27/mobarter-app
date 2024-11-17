package waitlist

import (
	"mobarter/app"

	"github.com/graphql-go/graphql"
)

var ResolverJoinWaitList = &graphql.Field{
	Type: graphql.NewObject(graphql.ObjectConfig{
		Name: "waitListResponse",
		Fields: graphql.Fields{
			"message": app.Str,
		},
	}),
	Args: graphql.FieldConfigArgument{
		"email":     app.ArgStr,
		"firstName": app.ArgStr,
		"lastName":  app.ArgStr,
		"country":   app.ArgStr,
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {

		// repo := NewRepository()

		return map[string]interface{}{
			"name": "Meat Pie",
		}, nil
	},
}
