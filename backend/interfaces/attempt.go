package interfaces

import s "techno/backend/structs"

type AttemptService interface {
	StartAttempt(userID, questionID int) (*s.Attempt, error)
	SubmitAttempt(attemptID int) error
	GetAttempt(attemptID int) (*s.Attempt, error)
}

type AttemptRepository interface {
	Create(attempt *s.Attempt) error
	FindByID(id int) (*s.Attempt, error)
	Update(attempt *s.Attempt) error
}
