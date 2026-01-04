package services

import (
	"errors"
	"time"

	i "techno/backend/interfaces"
	s "techno/backend/structs"
)

type EvaluationService struct {
	evalRepo    i.EvaluationRepository
	attemptRepo i.AttemptRepository // Add this
	openAI      i.OpenAIClient
}

func NewEvaluationService(evalRepo i.EvaluationRepository, attemptRepo i.AttemptRepository, openAI i.OpenAIClient) *EvaluationService {
	return &EvaluationService{
		evalRepo:    evalRepo,
		attemptRepo: attemptRepo,
		openAI:      openAI,
	}
}

func (st *EvaluationService) Evaluate(attemptID int) (*s.Evaluation, error) {
	attempt, err := st.attemptRepo.FindByID(attemptID)
	if err != nil {
		return nil, err
	}

	// Only Submitted attempts can be evaluated
	if attempt.Status != "SUBMITTED" {
		return nil, errors.New("only submitted attempts can be evaluated")
	}

	// Call OpenAI to evaluate
	score, feedback, err := st.openAI.EvaluateAnswer(attempt.QuestionID, attempt.Answer)
	if err != nil {
		return nil, err
	}

	// Create Evaluation object
	evaluation := &s.Evaluation{
		AttemptID: attemptID,
		Score:     score,
		Feedback:  feedback,
		CreatedAt: time.Now(),
	}

	// Save evaluation
	if err := st.evalRepo.Create(evaluation); err != nil {
		return nil, err
	}

	// Update attempt status to Evaluated
	attempt.Status = "EVALUATED"
	attempt.Version += 1
	if err := st.attemptRepo.Update(attempt); err != nil {
		return nil, err
	}

	return evaluation, nil
}
