package structs

import "time"

type AttemptStatus string

const (
	Started    AttemptStatus = "started"
	Clarifying AttemptStatus = "clarifying"
	Designing  AttemptStatus = "designing"
	Submitted  AttemptStatus = "submitted"
	Evaluated  AttemptStatus = "evaluated"
)

type Attempt struct {
	ID          int
	UserID      int
	QuestionID  int
	Status      AttemptStatus
	Answer      string
	Version     int
	StartedAt   time.Time
	SubmittedAt *time.Time
}
