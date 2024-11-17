package field

import "github.com/graphql-go/graphql"

var QueryFields = graphql.Fields{
	"BankAccount_Get":     GetBankAccount,
	"BankAccount_GetAll":  GetAllBankAccount,
	"User_GetInfo":        UserInfo,
	"CountryList":         CountryList,
	"Transactions_GetAll": GetAllTransactions,
	"Transactions_GetOne": GetOneTransactions,
	"Kyc_GetLevel":        Kyc_GetLevel,
}

// Definite mutations
var MutationsFields = graphql.Fields{
	"BankAccount_Create": CreateBankAccount,
	"PurchaseAirtime":    PurchaseAirtime,
	"PurchaseData":       PurchaseData,
	"JoinWaitList":       joinWaitList,
	"Auth_SentOtp":       Auth_SentOtp,
	"Kyc_VerifyNin":      Kyc_VerifyNin,
	"Kyc_VerifyBvn":      Kyc_VerifyBvn,
	"Kyc_VerifyEmail":    Kyc_VerifyEmail,
	"Kyc_VerifyPhone":    Kyc_VerifyPhone,
}
