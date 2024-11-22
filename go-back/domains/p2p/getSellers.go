package p2p

import (
	"errors"
	"mobarter/app"

	"github.com/graphql-go/graphql"
)

func GetP2PSellers(appState app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "GetP2PSellersResponse",
			Fields: graphql.Fields{
				"name":   app.String,
				"rate":   app.String,
				"amount": app.String,
			},
		}),
		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type: graphql.NewInputObject(
					graphql.InputObjectConfig{
						Name: "GetSellersInput",
						Fields: graphql.InputObjectConfigFieldMap{
							"paymentMethod": &graphql.InputObjectFieldConfig{
								Type: graphql.NewNonNull(graphql.Int),
							},
						},
					},
				),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// Extract arguments from params
			payMethod, ok := p.Args["paymentMethod"].(string)

			if !ok || payMethod == "" {
				return nil, errors.New("invalid or missing 'payMethod' argument")
			}

			// Example: Fetch the user data (mocked here)
			user := map[string]interface{}{
				"rate":   payMethod,
				"name":   "John Doe",
				"amount": "John Doe",
			}

			// Check if user exists (mock example)
			// if user["id"] == "" {
			// 	return nil, errors.New("user not found")
			// }

			return user, nil

		},
	}
}
