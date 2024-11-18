package database

import (
	"gorm.io/gorm"
)

// The order of which you run the migrations matters
// User first, workspace next
// Tables with More dependencies stays on top
func RunMigrations(db *gorm.DB) {
	err := db.AutoMigrate(
		User{},
		KycCredential{},
		Waitlist{},
		BankAccount{},
		AirtimeTransactions{},
		DataTransactions{},
		WalletTransactions{},
	)

	if err != nil {
		println("Could not run migrations, %v", err.Error())
	}
}
