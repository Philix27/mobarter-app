package field

import (
	"mobarter/app"
	"mobarter/field/bank"
	"mobarter/field/user"

	"github.com/graphql-go/graphql"
)

func QueryFields(appState app.AppState) graphql.Fields {
	return graphql.Fields{
		"BankAccount_Get":     bank.GetBankAccount(appState),
		"BankAccount_GetAll":  bank.GetAllBankAccount(appState),
		"User_GetInfo":        user.UserInfo(appState),
		"CountryList":         CountryList(appState),
		"Transactions_GetAll": GetAllTransactions(appState),
		"Transactions_GetOne": GetOneTransactions(appState),
		"Kyc_GetLevel":        Kyc_GetLevel(appState),
	}
}

func MutationsFields(appState app.AppState) graphql.Fields {
	return graphql.Fields{
		"BankAccount_Create": bank.CreateBankAccount(appState),
		"BankAccount_Delete": bank.DeleteBankAccount(appState),
		"User_Create":        user.Create(appState),
		"PurchaseAirtime":    PurchaseAirtime(appState),
		"PurchaseData":       PurchaseData(appState),
		"JoinWaitList":       JoinWaitList(appState),
		"Auth_SentOtp":       Auth_SentOtp(appState),
		"Kyc_VerifyNin":      Kyc_VerifyNin(appState),
		"Kyc_VerifyBvn":      Kyc_VerifyBvn(appState),
		"Kyc_VerifyEmail":    Kyc_VerifyEmail(appState),
		"Kyc_VerifyPhone":    Kyc_VerifyPhone(appState),
	}

}
