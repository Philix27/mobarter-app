package airtime

import (
	"errors"
	"mobarter/app"

	"github.com/LNMMusic/optional"
	"github.com/graphql-go/graphql"
)

type PurchaseDataDto struct {
	phone           string
	amount          int
	transactionHash string
	network         string
	walletAddress   string
	planId          string
}

func PurchaseData(appState app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "PurchaseDataResponse",
			Fields: graphql.Fields{
				"message": app.String,
			},
		}),
		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type: graphql.NewInputObject(
					graphql.InputObjectConfig{
						Name: "PurchaseDataInput",
						Fields: graphql.InputObjectConfigFieldMap{
							"amount":          app.InputInt,
							"phone":           app.InputStringOptional,
							"transactionHash": app.InputString,
							"network":         networkProvider,
							"walletAddress":   app.InputString,
							"planId":          app.InputString,
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

			dto := &PurchaseDataDto{
				amount:          args["amount"].(int),
				walletAddress:   args["walletAddress"].(string),
				phone:           args["phone"].(string),
				transactionHash: args["transactionHash"].(string),
				network:         args["network"].(string),
				planId:          args["planId"].(string),
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

			if !rechargePhoneWithData(dto.phone) {
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
