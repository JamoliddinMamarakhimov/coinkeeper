package repository

import (
	"coinkeeper/db"
	"coinkeeper/logger"
	"coinkeeper/models"
)

func GetAllIncomes(userID uint, query string) ([]models.Income, error) {
	var incomes []models.Income

	query = "%" + query + "%"

	err := db.GetDBConn().Model(&models.Income{}).
		Joins("Join users on users.id = incomes.user_id").
		Where("incomes.user_id = ? AND description iLIKE ?", userID, query).
		Order("incomes.id").
		Find(&incomes).Error
	if err != nil {
		logger.Error.Println("[repository.GetAllIncomes] cannot get all incomes. Error is:", err.Error())
		return nil, translateError(err)
	}
	return incomes, nil
}

func GetIncomeByID(userID, incomeID uint) (i models.Income, err error) {
	err = db.GetDBConn().Model(&models.Income{}).
		Joins("JOIN users ON users.id = incomes.user_id").
		Where("incomes.user_id = ? AND incomes.id = ?", userID, incomeID).
		First(&i).Error
	if err != nil {
		logger.Error.Println("[repository.GetIncomeByID] cannot get income by id. Error is:", err.Error())
		return models.Income{}, translateError(err)
	}
	return i, nil
}

func CreateIncome(i models.Income) error {
	err := db.GetDBConn().Create(&i).Error
	if err != nil {
		logger.Error.Println("[repository.CreateOperation] cannot create operation. Error is:", err.Error())
		return translateError(err)
	}
	return nil
}

func UpdateIncome(i models.Income) error {
	err := db.GetDBConn().Save(&i).Error
	if err != nil {
		logger.Error.Println("[repository.UpdateIncome] cannot update operation. Error is:", err.Error())
		return translateError(err)
	}
	return nil
}

func DeleteIncome(incomeID int, userID uint) error {
	err := db.GetDBConn().
		Table("incomes").
		Where("id = ? AND user_id = ?", incomeID, userID).
		Update("is_deleted", true).Error
	if err != nil {
		logger.Error.Println("[repository.DeleteIncome] cannot delete operation. Error is:", err.Error())
		return err
	}
	return nil
}
