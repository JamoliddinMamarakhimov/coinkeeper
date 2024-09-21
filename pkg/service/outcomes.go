package service

import (
	"coinkeeper/errs"
	"coinkeeper/models"
	"coinkeeper/pkg/repository"
	"errors"
)

func GetAllOutcomes(userID uint, query string) (outcomes []models.Outcome, err error) {
	outcomes, err = repository.GetAllOutcomes(userID, query)
	if err != nil {
		return nil, err
	}
	return outcomes, nil
}

func GetOutcomeByID(userID, outcomeID uint) (o models.Outcome, err error) {
	o, err = repository.GetOutcomeByID(userID, outcomeID)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return o, errs.ErrOperationNotFound
		}
		return models.Outcome{}, err
	}
	return o, nil
}

func CreateOutcome(o models.Outcome) error {
	if err := repository.CreateOutcome(&o); err != nil {
		return err
	}
	return nil
}

func UpdateOutcome(o models.Outcome) error {
	if err := repository.UpdateOutcome(&o); err != nil {
		return err
	}
	return nil
}

func DeleteOutcome(outcomeID int, userID uint) error {
	if err := repository.DeleteOutcome(outcomeID, userID); err != nil {
		return err
	}
	return nil
}
