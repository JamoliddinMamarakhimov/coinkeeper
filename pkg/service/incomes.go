package service

import (
	"coinkeeper/errs"
	"coinkeeper/models"
	"coinkeeper/pkg/repository"
	"errors"
)

func GetAllIncomes(userID uint, query string) (incomes []models.Income, err error) {
	incomes, err = repository.GetAllIncomes(userID, query)
	if err != nil {
		return nil, err
	}
	return incomes, nil
}

func GetIncomeByID(userID, incomeID uint) (i models.Income, err error) {
	i, err = repository.GetIncomeByID(userID, incomeID)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return i, errs.ErrOperationNotFound
		}
		return models.Income{}, err
	}
	return i, nil
}

func CreateIncome(i models.Income) error {
	if err := repository.CreateIncome(i); err != nil {
		return err
	}
	return nil
}

func UpdateIncome(i models.Income) error {
	if err := repository.UpdateIncome(i); err != nil {
		return err
	}
	return nil
}

func DeleteIncome(incomeID int, userID uint) error {
	if err := repository.DeleteIncome(incomeID, userID); err != nil {
		return err
	}
	return nil
}
