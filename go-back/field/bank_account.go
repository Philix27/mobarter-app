package field

import (
	"mobarter/app"

	"github.com/graphql-go/graphql"
)

func GetBankAccount(appState app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "GetBankAccountResponse",
			Fields: graphql.Fields{
				"name":           String,
				"id":             Int,
				"dexId":          Int,
				"upVotes":        Int,
				"downVotes":      Int,
				"totalVotes":     Int,
				"winPercentage":  Float,
				"lossPercentage": Float,
			},
		}),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			return map[string]interface{}{
				"name": "Meat Pie",
			}, nil
		},
	}
}

func CreateBankAccount(appState app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "CreateBankAccountResponse",
			Fields: graphql.Fields{
				"message": String,
			},
		}),
		Args: graphql.FieldConfigArgument{
			"accountName": argString,
			"accountNo":   argString,
			"bankName":    argString,
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			return map[string]interface{}{
				"name": "Meat Pie",
			}, nil
		},
	}
}

func GetAllBankAccount(appState app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "getAllBankAccountResponse",
			Fields: graphql.Fields{
				"accountName": String,
				"accountNo":   String,
				"bankName":    String,
			},
		}),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			return map[string]interface{}{
				"name": "Meat Pie",
			}, nil
		},
	}
}
