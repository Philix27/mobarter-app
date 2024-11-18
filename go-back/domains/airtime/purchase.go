package airtime

import (
	"errors"
	"mobarter/app"
	"mobarter/log"

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
			"issuerAddress":   app.ArgString,
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			ilog := log.New("Purchase airtime")

			dto := &PurchaseAirtimeDto{
				amount:          p.Args["amount"].(int),
				issuerAddress:   p.Args["issuerAddress"].(string),
				phone:           p.Args["phone"].(string),
				transactionHash: p.Args["transactionHash"].(string),
				network:         p.Args["network"].(string),
			}
			// todo: Verify the amount is greater than 0

			if dto.amount < 100 {
				return map[string]interface{}{
					"message": "Amount is too small",
				}, errors.New("Amount is too small")
			}
			// todo: Validate the network
			if !verifyNetwork(dto.network) {
				return map[string]interface{}{
					"message": "Invalid network provider",
				}, errors.New("Invalid network providerl")
			}
			// todo: Verify the transaction hash
			if !verifyTransactionHash(dto.transactionHash) {
				return map[string]interface{}{
					"message": "Could not verify the transaction",
				}, nil
			}
			// todo: Send request to Flutterwave and await response
			if !rechargePhone(dto.phone) {
				return map[string]interface{}{
					"message": "Could not be recharged",
				}, nil
			}
			// todo: parse response to struct

			// todo: add to transactions for email
			addTransaction(dto.phone, app.TransactionAirtime)
			// todo: log
			ilog.Debug("Recharged airtime for" + dto.phone)
			// todo: send response
			return map[string]interface{}{
				"name": "Meat Pie",
			}, nil
		},
	}
}
