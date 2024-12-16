package field

import (
	"mobarter/app"
	"mobarter/domains/admin"
	"mobarter/domains/airtime"
	"mobarter/domains/auth"
	"mobarter/domains/bank"
	"mobarter/domains/kyc"
	"mobarter/domains/p2p"
	"mobarter/domains/user"

	"github.com/graphql-go/graphql"
)

func QueryFields(appState app.AppState) graphql.Fields {
	return graphql.Fields{
		"BankAccount_Get":      bank.GetBankAccount(appState),
		"BankAccount_GetAll":   bank.GetAllBankAccount(appState),
		"User_GetInfo":         user.UserInfo(appState),
		"GetBuyers":            p2p.GetBuyers(appState),
		"GetP2PSellers":        p2p.GetP2PSellers(appState),
		"CountryList":          CountryList(appState),
		"Transactions_GetAll":  GetAllTransactions(appState),
		"Transactions_GetOne":  GetOneTransactions(appState),
		"Kyc_GetLevel":         kyc.Kyc_GetLevel(appState),
		"Airtime_GetDataPlans": airtime.GetDataPlans(appState),
		//! Admin
		"Admin_GetBlockedAccounts":  admin.GetBlockedAccounts(appState),
		"Admin_GetUserBankAccounts": admin.GetUserBankAccounts(appState),
		"Admin_GetUserTransactions": admin.GetUserTransactions(appState),
		"Admin_GetUsers":            admin.GetUsers(appState),
	}
}

func MutationsFields(appState app.AppState) graphql.Fields {
	return graphql.Fields{
		"JoinWaitList": JoinWaitList(appState),
		"Auth_SentOtp": Auth_SentOtp(appState),
		//! bank
		"BankAccount_Create": bank.CreateBankAccount(appState),
		"BankAccount_Delete": bank.DeleteBankAccount(appState),
		//! Purchase Data & airtime
		"PurchaseAirtime": airtime.PurchaseAirtime(appState),
		"PurchaseData":    airtime.PurchaseData(appState),
		//! kyc
		"Kyc_VerifyNin":   kyc.Kyc_VerifyNin(appState),
		"Kyc_VerifyBvn":   kyc.Kyc_VerifyBvn(appState),
		"Kyc_VerifyEmail": kyc.Kyc_VerifyEmail(appState),
		"Kyc_VerifyPhone": kyc.Kyc_VerifyPhone(appState),
		//! Admin
		"Admin_ApproveBvn":                admin.ApproveBvn(appState),
		"Admin_ApproveNin":                admin.ApproveNin(appState),
		"Admin_ApprovePassport":           admin.ApprovePassport(appState),
		"Admin_ApproveResidentialAddress": admin.ApproveResidentialAddress(appState),
		"Admin_BlockAccount":              admin.BlockAccount(appState),
		"Admin_Login":                     admin.Login(appState),
		"Admin_SendEmail":                 admin.SendEmail(appState),
		"Admin_SetRates":                  admin.SetRates(appState),
		// ! Auth
		"Auth_SendEmailOtp": auth.Auth_SendEmailOtp(appState),
		"Auth_SendPhoneOtp": auth.Auth_SendPhoneOtp(appState),
		"Auth_VerifyOtp":    auth.Auth_VerifyOtp(appState),
		// ! User
		"User_Create":         user.Create(appState),
		"User_ResetPassword":  user.ResetPassword(appState),
		"User_AddNin":         user.AddNin(appState),
		"User_AddBvn":         user.AddBvn(appState),
		"User_PhoneSendOtp":   user.PhoneSendOtp(appState),
		"User_PhoneVerifyOtp": user.PhoneVerifyOtp(appState),
	}

}
