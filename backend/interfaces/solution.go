package interfaces

import s "techno/backend/structs"

type SolutionRepository interface {
	Create(solution *s.Solution) error
	FindByAttemptID(attemptID int) (*s.Solution, error)
}
