package user

import (
	"mobarter/app"
	"mobarter/database"
)

func createUserRepo(appState app.AppState, dto *CreateDto) error {
	result := appState.DB.Create(&database.User{
		WalletAddress: dto.WalletAddress,
		FirstName:     dto.FirstName,
		LastName:      dto.LastName,
	})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func findByWalletAddress(appState app.AppState, walletAddress string) (database.User, error) {
	obj := database.User{}

	result := appState.DB.Where("wallet_address = ?", walletAddress).First(&obj)

	if result.Error != nil {
		// r.logger.Error("FindById:", result.Error)
		return database.User{}, result.Error
	} else {
		// r.logger.Error("RETRIEVE_FIND_ONE:", object)
		return obj, nil
	}

}
