package airtime

import (
	"errors"
	"mobarter/app"

	"github.com/graphql-go/graphql"
)

type GetDataPlansDto struct {
	category string
	network  string
}

func GetDataPlans(appState app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "GetDataPlansResponse",
			Fields: graphql.Fields{
				"message": app.String,
			},
		}),
		Args: graphql.FieldConfigArgument{
			"network":  app.ArgString,
			"category": app.ArgOptionalString,
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// ilog := log.New("Purchase airtime")

			dto := &GetDataPlansDto{
				network:  p.Args["network"].(string),
				category: p.Args["category"].(string),
			}

			if !verifyNetwork(dto.network) {
				return nil, errors.New("Invalid network provider")
			}

			// todo: log
			// ilog.Debug("Recharged airtime for" + dto.phone)

			return map[string]interface{}{
				"message": "success",
			}, nil
		},
	}
}
