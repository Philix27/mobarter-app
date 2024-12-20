package advert

import (
	"mobarter/app"

	"github.com/graphql-go/graphql"
)

var supportedCountries = &graphql.InputObjectFieldConfig{
	Type: graphql.NewEnum(
		graphql.EnumConfig{
			Name: "AdvertCountries",
			Values: graphql.EnumValueConfigMap{
				"NIGERIA": &graphql.EnumValueConfig{
					Value: "NIGERIA",
				},
				"KENYA": &graphql.EnumValueConfig{
					Value: "KENYA",
				},
				"GHANA": &graphql.EnumValueConfig{
					Value: "GHANA",
				},
			},
		},
	),
}
var supportedPaymentMethods = &graphql.InputObjectFieldConfig{
	Type: graphql.NewEnum(
		graphql.EnumConfig{
			Name: "AdvertPaymentMethods",
			Values: graphql.EnumValueConfigMap{
				"BANK_ACCOUNT": &graphql.EnumValueConfig{
					Value: "BANK_ACCOUNT",
				},
				"TRANSFER": &graphql.EnumValueConfig{
					Value: "TRANSFER",
				},
				"MOBILE_MONEY": &graphql.EnumValueConfig{
					Value: "MOBILE_MONEY",
				},
			},
		},
	),
}

func GetAllAdvert(appState app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "Advert_GetAllResponse",
			Fields: graphql.Fields{
				"message":       app.String,
				"id":            app.String,
				"rate":          app.String,
				"limit_lower":   app.String,
				"limit_upper":   app.String,
				"currency_pair": app.String,
				"instructions":  app.String,
			},
		}),

		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type: graphql.NewInputObject(
					graphql.InputObjectConfig{
						Name: "Advert_GetAllInput",
						Fields: graphql.InputObjectConfigFieldMap{
							"token":         app.InputString,
							"country":       supportedCountries,
							"paymentMethod": supportedPaymentMethods,
						},
					},
				),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// ilog := log.New("Get user info")
			//

			return map[string]interface{}{
				"message": "success",
			}, nil
		},
	}
}
