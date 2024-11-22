package airtime

import (
	"mobarter/app"

	"github.com/graphql-go/graphql"
)

var dataPlans = []map[string]interface{}{
	{"id": "m1", "network": "MTN", "plan": "1gb #500 - 2days"},
	{"id": "m2", "network": "MTN", "plan": "1.5gb #500 - 3days"},
	{"id": "m3", "network": "MTN", "plan": "1gb #400 - 6days"},
	{"id": "m4", "network": "MTN", "plan": "1gb #400 - 2 weeks"},
	{"id": "m5", "network": "MTN", "plan": "1gb #400 - 10 days"},
	{"id": "a1", "network": "AIRTEL", "plan": "1gb #500 - 2days"},
	{"id": "a2", "network": "AIRTEL", "plan": "1gb #2500 - 2days"},
	{"id": "a3", "network": "AIRTEL", "plan": "1gb #1500 - 2days"},
	{"id": "a4", "network": "AIRTEL", "plan": "1gb #2000 - 2days"},
	{"id": "a5", "network": "AIRTEL", "plan": "1gb #350 - 2days"},
	{"id": "g1", "network": "GLO", "plan": "1gb #1500 - 1 week"},
	{"id": "g2", "network": "GLO", "plan": "1gb #3500 - 1 week"},
	{"id": "g3", "network": "GLO", "plan": "1gb #7500 - 1 week"},
	{"id": "g4", "network": "GLO", "plan": "1gb #7500 - 1 week"},
}

type GetDataPlansDto struct {
	category string
	network  string
}

func GetDataPlans(appState app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "GetDataPlansResponse",
			Fields: graphql.Fields{
				"dataPlans": &graphql.Field{
					Type: graphql.NewList(
						graphql.NewObject(
							graphql.ObjectConfig{
								Name: "DataPlans",
								Fields: graphql.Fields{
									"id":      app.String,
									"network": app.String,
									"plan": app.String,
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
						Name: "GetDataPlansInput",
						Fields: graphql.InputObjectConfigFieldMap{
							"network":  networkProvider,
							"category": app.InputStringOptional,
						},
					},
				),
			},
		},

		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// ilog := log.New("Purchase airtime")
			// args, err := app.ValidateArg(optional.Some("input"), p)
			// if err != nil {
			// 	return nil, errors.New("Input required")
			// }
			// dto := &GetDataPlansDto{
			// 	network:  args["network"].(string),
			// 	category: args["category"].(string),
			// }

			// if !verifyNetwork(dto.network) {
			// 	return nil, errors.New("Invalid network provider")
			// }

			// todo: log
			// ilog.Debug("Recharged airtime for" + dto.phone)

			return map[string]interface{}{
				"dataPlans": dataPlans,
			}, nil
		},
	}
}
