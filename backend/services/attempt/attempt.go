package services

import (
	"errors"
	i "techno/backend/interfaces"
	s "techno/backend/structs"
	"time"
)

type AttemptService struct {
	repo i.AttemptRepository
}

func NewAttemptService(repo i.AttemptRepository) *AttemptService {
	return &AttemptService{repo: repo}
}

func (a *AttemptService) StartAttempt(userID, questionID int) (*s.Attempt, error) {
	if userID == 0 || questionID == 0 {
		return nil, errors.New("userID and questionID required")
	}

	attempt := &s.Attempt{
		UserID:     userID,
		QuestionID: questionID,
		Status:     "STARTED",
		StartedAt:  time.Now(),
	}

	if err := a.repo.Create(attempt); err != nil {
		return nil, err
	}

	return attempt, nil
}

func (a *AttemptService) SubmitAttempt(attemptID int) error {
	if attemptID == 0 {
		return errors.New("attemptID required")
	}

	attempt, err := a.repo.FindByID(attemptID)
	if err != nil {
		return err
	}

	if attempt.Status != "STARTED" {
		return errors.New("attempt already submitted")
	}

	now := time.Now()
	attempt.Status = "SUBMITTED"
	attempt.SubmittedAt = &now

	return a.repo.Update(attempt)
}
