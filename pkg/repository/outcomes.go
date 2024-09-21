package repository

import (
	"coinkeeper/db"
	"coinkeeper/logger"
	"coinkeeper/models"
)

func GetAllOutcomes(userID uint, query string) ([]models.Outcome, error) {
	var outcomes []models.Outcome

	query = "%" + query + "%"

	err := db.GetDBConn().Model(&models.Outcome{}).
		Joins("JOIN users ON users.id = outcomes.user_id").
		Where("outcomes.user_id = ? AND description ILIKE ?", userID, query).
		Order("outcomes.id").
		Find(&outcomes).Error
	if err != nil {
		logger.Error.Println("[repository.GetAllOutcomes] cannot get all outcomes. Error is:", err.Error())
		return nil, translateError(err)
	}
	return outcomes, nil
}

func GetOutcomeByID(userID, outcomeID uint) (models.Outcome, error) {
	var outcome models.Outcome

	err := db.GetDBConn().Model(&models.Outcome{}).
		Joins("JOIN users ON users.id = outcomes.user_id").
		Where("outcomes.user_id = ? AND outcomes.id = ?", userID, outcomeID).
		First(&outcome).Error
	if err != nil {
		logger.Error.Println("[repository.GetOutcomeByID] cannot get outcome by id. Error is:", err.Error())
		return models.Outcome{}, translateError(err)
	}
	return outcome, nil
}

func CreateOutcome(outcome *models.Outcome) error {
	err := db.GetDBConn().Create(outcome).Error
	if err != nil {
		logger.Error.Println("[repository.CreateOutcome] cannot create outcome. Error is:", err.Error())
		return translateError(err)
	}
	return nil
}

func UpdateOutcome(outcome *models.Outcome) error {
	err := db.GetDBConn().Save(outcome).Error
	if err != nil {
		logger.Error.Println("[repository.UpdateOutcome] cannot update outcome. Error is:", err.Error())
		return translateError(err)
	}
	return nil
}

func DeleteOutcome(outcomeID int, userID uint) error {
	err := db.GetDBConn().
		Table("outcomes").
		Where("id = ? AND user_id = ?", outcomeID, userID).
		Update("is_deleted", true).Error
	if err != nil {
		logger.Error.Println("[repository.DeleteOutcome] cannot delete outcome. Error is:", err.Error())
		return err
	}
	return nil
}
