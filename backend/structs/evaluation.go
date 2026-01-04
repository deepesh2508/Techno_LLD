package structs

import "time"

type Evaluation struct {
	ID          int
	CreatedAt   time.Time
	AttemptID   int
	Score       float64
	Feedback    string
	EvaluatedBy string
}
