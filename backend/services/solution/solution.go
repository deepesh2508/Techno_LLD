package services

import (
	"errors"

	i "techno/backend/interfaces"
	s "techno/backend/structs"
)

type SolutionService struct {
	repo i.SolutionRepository
}

func NewSolutionService(repo i.SolutionRepository) *SolutionService {
	return &SolutionService{repo: repo}
}

func (s *SolutionService) Submit(solution s.Solution) error {
	// 1️⃣ Basic validation: AttemptID required
	if solution.AttemptID == 0 {
		return errors.New("AttemptID is required")
	}

	// Optional: Validate at least FunctionalReq + Code is present
	if solution.FunctionalReq == "" && solution.Code == "" {
		return errors.New("either FunctionalReq or Code must be provided")
	}

	// 2️⃣ Save to repository
	if err := s.repo.Create(&solution); err != nil {
		return err
	}

	return nil
}
