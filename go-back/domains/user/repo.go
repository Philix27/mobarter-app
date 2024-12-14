package user

import (
	"errors"
	"mobarter/app"
	"mobarter/db"
)

func createUserRepo(appState app.AppState, dto *CreateDto) error {

	return nil
}

func findByWalletAddress(appState app.AppState, walletAddress string) (db.User, error) {
	obj := db.User{}

	// result := appState.DB.Where("wallet_address = ?", walletAddress).First(&obj)

	// if result.Error != nil {
	// 	// r.logger.Error("FindById:", result.Error)
	// 	return database.User{}, result.Error
	// } else {
	// 	// r.logger.Error("EVE_FIND_ONE:", object)
	// 	return obj, nil
	// }
	return obj, errors.New("")
}
