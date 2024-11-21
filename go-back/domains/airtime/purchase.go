package airtime

import (
	"errors"
	"mobarter/app"

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
			"phone":           app.ArgString,
			"amount":          app.ArgInt,
			"transactionHash": app.ArgString,
			"network":         app.ArgString,
			"walletAddress":   app.ArgString,
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// ilog := log.New("Purchase airtime")

			dto := &PurchaseAirtimeDto{
				amount:          p.Args["amount"].(int),
				issuerAddress:   p.Args["walletAddress"].(string),
				phone:           p.Args["phone"].(string),
				transactionHash: p.Args["transactionHash"].(string),
				network:         p.Args["network"].(string),
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
