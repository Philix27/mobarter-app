package database

import (
	"gorm.io/gorm"
)

// The order of which you run the migrations matters
// User first, workspace next
// Tables with More dependencies stays on top
func RunMigrations(db *gorm.DB) {
	err := db.AutoMigrate(
		KycCredential{},
		Waitlist{},
		BankAccount{},
		AirtimeTransactions{},
		DataTransactions{},
		WalletTransactions{},
		User{},
	)

	if err != nil {
		println("Could not run migrations, %v", err.Error())
	}
}
