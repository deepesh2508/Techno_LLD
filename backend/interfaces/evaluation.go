package interfaces

import s "techno/backend/structs"

type EvaluationService interface {
	Evaluate(attemptID string) (*s.Evaluation, error)
}

type EvaluationRepository interface {
	Create(e *s.Evaluation) error
}
