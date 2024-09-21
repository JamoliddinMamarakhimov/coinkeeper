package service

import (
	"coinkeeper/errs"
	"coinkeeper/models"
	"coinkeeper/pkg/repository"
	"errors"
)

func GetAllAccounts(userID uint, query string) (accounts []models.Account, err error) {
	accounts, err = repository.GetAllAccounts(userID, query)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func GetAccountByID(userID, accountID uint) (a models.Account, err error) {
	a, err = repository.GetAccountByID(userID, accountID)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return a, errs.ErrOperationNotFound
		}
		return models.Account{}, err
	}
	return a, nil
}

func CreateAccount(a models.Account) error {
	if err := repository.CreateAccount(a); err != nil {
		return err
	}
	return nil
}

func UpdateAccount(a models.Account) error {
	if err := repository.UpdateAccount(a); err != nil {
		return err
	}
	return nil
}

func DeleteAccount(accountID uint, userID uint) error {
	if err := repository.DeleteAccount(accountID, userID); err != nil {
		return err
	}
	return nil
}
