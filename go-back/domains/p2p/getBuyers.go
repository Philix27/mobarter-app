package p2p

import (
	"errors"
	"mobarter/app"

	"github.com/graphql-go/graphql"
)

func GetBuyers(appState app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "GetP2PBuyersResponse",
			Fields: graphql.Fields{
				"accountName": app.String,
				"accountNo":   app.String,
				"bankName":    app.String,
			},
		}),
		Args: graphql.FieldConfigArgument{
			"paymentMethod": app.ArgString,
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// Extract arguments from params
			id, ok := p.Args["id"].(string)
			if !ok || id == "" {
				return nil, errors.New("invalid or missing 'id' argument")
			}

			// Example: Fetch the user data (mocked here)
			user := map[string]interface{}{
				"id":   id,
				"name": "John Doe",
			}

			// Check if user exists (mock example)
			if user["id"] == "" {
				return nil, errors.New("user not found")
			}

			return user, nil

		},
	}
}
