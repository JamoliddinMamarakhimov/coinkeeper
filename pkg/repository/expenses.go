package repository

import (
	"coinkeeper/db"
	"coinkeeper/logger"
	"coinkeeper/models"
)

func GetAllExpenses(userID uint, query string) ([]models.Expense, error) {
	var expenses []models.Expense

	query = "%" + query + "%"

	err := db.GetDBConn().Model(&models.Expense{}).
		Joins("JOIN users ON users.id = expenses.user_id").
		Where("expenses.user_id = ? AND description ILIKE ?", userID, query).
		Order("expenses.id").
		Find(&expenses).Error
	if err != nil {
		logger.Error.Println("[repository.GetAllExpenses] cannot get all expenses. Error is:", err.Error())
		return nil, translateError(err)
	}
	return expenses, nil
}

func GetExpenseByID(userID, expenseID uint) (models.Expense, error) {
	var expense models.Expense

	err := db.GetDBConn().Model(&models.Expense{}).
		Joins("JOIN users ON users.id = expenses.user_id").
		Where("expenses.user_id = ? AND expenses.id = ?", userID, expenseID).
		First(&expense).Error
	if err != nil {
		logger.Error.Println("[repository.GetExpenseByID] cannot get expense by id. Error is:", err.Error())
		return models.Expense{}, translateError(err)
	}
	return expense, nil
}

func CreateExpense(expense *models.Expense) error {
	err := db.GetDBConn().Create(expense).Error
	if err != nil {
		logger.Error.Println("[repository.CreateExpense] cannot create expense. Error is:", err.Error())
		return translateError(err)
	}
	return nil
}

func UpdateExpense(expense *models.Expense) error {
	err := db.GetDBConn().Save(expense).Error
	if err != nil {
		logger.Error.Println("[repository.UpdateExpense] cannot update expense. Error is:", err.Error())
		return translateError(err)
	}
	return nil
}

func DeleteExpense(expenseID, userID uint) error {
	err := db.GetDBConn().
		Table("expenses").
		Where("id = ? AND user_id = ?", expenseID, userID).
		Update("is_deleted", true).Error
	if err != nil {
		logger.Error.Println("[repository.DeleteExpense] cannot delete expense. Error is:", err.Error())
		return err
	}
	return nil
}
