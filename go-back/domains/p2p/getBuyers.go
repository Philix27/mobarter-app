package p2p

import (
	"mobarter/app"

	"github.com/graphql-go/graphql"
)

var VendorType = &graphql.InputObjectFieldConfig{
	Type: graphql.NewEnum(
		graphql.EnumConfig{
			Name: "VendorType",
			Values: graphql.EnumValueConfigMap{
				"BUYERS": &graphql.EnumValueConfig{
					Value: "BUYERS",
				},
				"SELLERS": &graphql.EnumValueConfig{
					Value: "SELLERS",
				},
			},
		},
	),
}

var dataPlans = []map[string]interface{}{
	{"id": "1", "name": "John Doe"},
	{"id": "2", "name": "Jane Smith"},
	{"id": "3", "name": "Alice Johnson"},
}

func GetBuyers(appState app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "GetP2PBuyersResponse",
			Fields: graphql.Fields{
				"accountName": app.String,
				"accountNo":   app.String,
				"bankName":    app.String,
				"vendors": &graphql.Field{
					Type: graphql.NewList(
						graphql.NewObject(
							graphql.ObjectConfig{
								Name: "User",
								Fields: graphql.Fields{
									"id":   app.String,
									"name": app.String,
								},
							},
						),
					),
				},
			},
		}),
		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type: graphql.NewInputObject(
					graphql.InputObjectConfig{
						Name: "GetBuyersInput",
						Fields: graphql.InputObjectConfigFieldMap{
							"vendor": VendorType,
						},
					},
				),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// Extract arguments from params
			// id, ok := p.Args["id"].(string)
			// if !ok || id == "" {
			// 	return nil, errors.New("invalid or missing 'id' argument")
			// }

			// Check if user exists (mock example)
			// if user["id"] == "" {
			// 	return nil, errors.New("user not found")
			// }

			return map[string]interface{}{
				"plans": dataPlans,
				// "accountName": "John Doe",
				// "accountNo":   "John Doe",
				// "bankName":    "John Doe",
			}, nil

		},
	}
}
