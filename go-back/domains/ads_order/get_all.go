package orders

import (
	"mobarter/app"

	"github.com/graphql-go/graphql"
)

var orderStatus = &graphql.InputObjectFieldConfig{
	Type: graphql.NewEnum(
		graphql.EnumConfig{
			Name: "OrderStatus",
			Values: graphql.EnumValueConfigMap{
				"COMPLETED": &graphql.EnumValueConfig{
					Value: "COMPLETED",
				},
				"PENDING": &graphql.EnumValueConfig{
					Value: "PENDING",
				},
				"APPEAL": &graphql.EnumValueConfig{
					Value: "APPEAL",
				},
			},
		},
	),
}

func GetAll(appState app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "Order_GetAllResponse",
			Fields: graphql.Fields{
				"message": app.String,
			},
		}),
		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type: graphql.NewInputObject(
					graphql.InputObjectConfig{
						Name: "Order_GetAllInput",
						Fields: graphql.InputObjectConfigFieldMap{
							"status": orderStatus,
						},
					},
				),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			return map[string]interface{}{
				"message": "success",
			}, nil
		},
	}
}
