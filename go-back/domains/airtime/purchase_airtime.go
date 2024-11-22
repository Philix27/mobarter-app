package airtime

import (
	"errors"
	"mobarter/app"

	"github.com/LNMMusic/optional"
	"github.com/graphql-go/graphql"
)

type PurchaseAirtimeDto struct {
	phone           string
	amount          int
	transactionHash string
	network         string
	issuerAddress   string
}

func PurchaseAirtime(appState app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "PurchaseAirtimeResponse",
			Fields: graphql.Fields{
				"message": app.String,
			},
		}),
		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type: graphql.NewInputObject(
					graphql.InputObjectConfig{
						Name: "PurchaseAirtimeInput",
						Fields: graphql.InputObjectConfigFieldMap{
							"amount": &graphql.InputObjectFieldConfig{
								Type: graphql.NewNonNull(graphql.Int),
							},
							"phone": &graphql.InputObjectFieldConfig{
								Type: graphql.String,
							},
							"transactionHash": &graphql.InputObjectFieldConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"network": &graphql.InputObjectFieldConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"walletAddress": &graphql.InputObjectFieldConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
						},
					},
				),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// ilog := log.New("Purchase airtime")
			args, err := app.ValidateArg(optional.Some("input"), p)
			if err != nil {
				return nil, errors.New("Input required")
			}

			dto := &PurchaseAirtimeDto{
				amount:          args["amount"].(int),
				issuerAddress:   args["walletAddress"].(string),
				phone:           args["phone"].(string),
				transactionHash: args["transactionHash"].(string),
				network:         args["network"].(string),
			}

			if dto.amount < 100 {
				return nil, errors.New("Amount is too small")
			}

			if !verifyNetwork(dto.network) {
				return nil, errors.New("Invalid network provider")
			}

			if !verifyTransactionHash(dto.transactionHash) {
				return nil, errors.New("Could not verify the transaction")
			}

			if !rechargePhone(dto.phone) {
				return nil, errors.New("Phone number could not be recharged")
			}

			addTransaction(dto.phone, app.TransactionAirtime)
			// todo: log
			// ilog.Debug("Recharged airtime for" + dto.phone)

			return map[string]interface{}{
				"message": "success",
			}, nil
		},
	}
}
