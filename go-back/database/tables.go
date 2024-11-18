package database

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	WalletAddress string
	FirstName     string
	LastName      string
	BankAccounts  []BankAccount
	KycCredential KycCredential
}

type Waitlist struct {
	gorm.Model
	Email     string
	FirstName string
	LastName  string
}
type KycCredential struct {
	gorm.Model
	UserID                uint
	Email                 string
	IsEmailVerified       bool
	Phone                 string
	IsPhoneVerified       bool
	Nin                   string
	IsNinVerified         bool
	Bvn                   string
	IsBvnVerified         bool
	HomeAddress           string
	IsHomeAddressVerified bool
	NextOfKinName         string
	NextOfKinPhone        string
}

type BankAccount struct {
	gorm.Model
	Name     string
	No       string
	BankName string
	UserID   uint
}

// #region transactions
type WalletTransactions struct {
	gorm.Model
	Title  string
	Amount int
	From   string
	To     string
	Status string
}

type AirtimeTransactions struct {
	gorm.Model
	Phone   string
	Network string
	Amount  int
	Fee     int
	Status  string
	To      string
}
type DataTransactions struct {
	gorm.Model
	Phone      string
	Network    string
	Amount     int
	Fee        int
	Status     string
	DataPlanId string
	To         string
}

//#region
