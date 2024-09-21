package repository

import (
	"coinkeeper/db"
	"coinkeeper/logger"
	"coinkeeper/models"
)

func GetAllAccounts(userID uint, query string) ([]models.Account, error) {
	var accounts []models.Account
	query = "%" + query + "%"

	err := db.GetDBConn().Model(&models.Account{}).
		Joins("JOIN users ON users.id = accounts.user_id").
		Where("accounts.user_id = ? AND (accounts.card_number ILIKE ? OR accounts.description ILIKE ?)", userID, query, query).
		Order("accounts.id").
		Find(&accounts).Error
	if err != nil {
		logger.Error.Println("[repository.GetAllAccounts] cannot get all accounts. Error is:", err.Error())
		return nil, translateError(err)
	}
	return accounts, nil
}

func GetAccountByID(userID, accountID uint) (a models.Account, err error) {
	err = db.GetDBConn().Model(&models.Account{}).
		Joins("JOIN users ON users.id = accounts.user_id").
		Where("accounts.user_id = ? AND accounts.id = ?", userID, accountID).
		First(&a).Error
	if err != nil {
		logger.Error.Println("[repository.GetAccountByID] cannot get account by id. Error is:", err.Error())
		return models.Account{}, translateError(err)
	}
	return a, nil
}

func CreateAccount(a models.Account) error {
	err := db.GetDBConn().Create(&a).Error
	if err != nil {
		logger.Error.Println("[repository.CreateAccount] cannot create account. Error is:", err.Error())
		return translateError(err)
	}
	return nil
}

func UpdateAccount(a models.Account) error {
	err := db.GetDBConn().Save(&a).Error
	if err != nil {
		logger.Error.Println("[repository.UpdateAccount] cannot update account. Error is:", err.Error())
		return translateError(err)
	}
	return nil
}

func DeleteAccount(accountID uint, userID uint) error {
	err := db.GetDBConn().
		Table("accounts").
		Where("id = ? AND user_id = ?", accountID, userID).
		Update("is_deleted", true).Error
	if err != nil {
		logger.Error.Println("[repository.DeleteAccount] cannot delete account. Error is:", err.Error())
		return err
	}
	return nil
}
