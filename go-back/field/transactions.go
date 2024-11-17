package field

import "github.com/graphql-go/graphql"

var getAllTransactionsResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "getAllTransactionsResponse",
	Fields: graphql.Fields{
		"walletAddress": String,
		"firstName":     String,
		"lastName":      String,
		"country":       String,
	},
})

var GetAllTransactions *graphql.Field = &graphql.Field{
	Type: getAllTransactionsResponse,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {

		return map[string]interface{}{
			"name": "Meat Pie",
		}, nil
	},
}

var getOneTransactionsResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "getOneTransactionsResponse",
	Fields: graphql.Fields{
		"walletAddress": String,
		"firstName":     String,
		"lastName":      String,
		"country":       String,
	},
})

var GetOneTransactions *graphql.Field = &graphql.Field{
	Type: getAllTransactionsResponse,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {

		return map[string]interface{}{
			"name": "Meat Pie",
		}, nil
	},
}
