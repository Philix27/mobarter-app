package airtime

import (
	"errors"
	"mobarter/app"

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
			"phone":           app.ArgString,
			"amount":          app.ArgInt,
			"transactionHash": app.ArgString,
			"network":         app.ArgString,
			"walletAddress":   app.ArgString,
			"planId":          app.ArgString,
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// ilog := log.New("Purchase airtime")

			dto := &PurchaseDataDto{
				amount:          p.Args["amount"].(int),
				walletAddress:   p.Args["walletAddress"].(string),
				phone:           p.Args["phone"].(string),
				transactionHash: p.Args["transactionHash"].(string),
				network:         p.Args["network"].(string),
				planId:          p.Args["planId"].(string),
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
