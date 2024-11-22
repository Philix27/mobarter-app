package field

import (
	"mobarter/app"
	"mobarter/domains/airtime"
	"mobarter/domains/bank"
	"mobarter/domains/kyc"
	"mobarter/domains/p2p"
	"mobarter/domains/user"

	"github.com/graphql-go/graphql"
)

func QueryFields(appState app.AppState) graphql.Fields {
	return graphql.Fields{
		"BankAccount_Get":     bank.GetBankAccount(appState),
		"BankAccount_GetAll":  bank.GetAllBankAccount(appState),
		"User_GetInfo":        user.UserInfo(appState),
		"GetBuyers":           p2p.GetBuyers(appState),
		"GetP2PSellers":       p2p.GetP2PSellers(appState),
		"CountryList":         CountryList(appState),
		"Transactions_GetAll": GetAllTransactions(appState),
		"Transactions_GetOne": GetOneTransactions(appState),
		"Kyc_GetLevel":        kyc.Kyc_GetLevel(appState),
	}
}

func MutationsFields(appState app.AppState) graphql.Fields {
	return graphql.Fields{
		"BankAccount_Create": bank.CreateBankAccount(appState),
		"BankAccount_Delete": bank.DeleteBankAccount(appState),
		"User_Create":        user.Create(appState),
		"PurchaseAirtime":    airtime.PurchaseAirtime(appState),
		"PurchaseData":       airtime.PurchaseData(appState),
		"GetDataPlans":       airtime.GetDataPlans(appState),
		"Kyc_VerifyNin":      kyc.Kyc_VerifyNin(appState),
		"Kyc_VerifyBvn":      kyc.Kyc_VerifyBvn(appState),
		"Kyc_VerifyEmail":    kyc.Kyc_VerifyEmail(appState),
		"Kyc_VerifyPhone":    kyc.Kyc_VerifyPhone(appState),
		"JoinWaitList":       JoinWaitList(appState),
		// "Auth_SentOtp":       Auth_SentOtp(appState),
	}

}
