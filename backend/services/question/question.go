package question

import (
	"errors"
	"techno/backend/structs"

	"gorm.io/gorm"
)

type QuestionService struct {
	db *gorm.DB
}

func NewQuestionService(db *gorm.DB) *QuestionService {
	return &QuestionService{db: db}
}

func (s *QuestionService) Create(q *structs.Question) error {
	if q.Title == "" || q.Difficulty == "" {
		return errors.New("invalid question data")
	}
	return s.db.Create(q).Error
}

func (s *QuestionService) GetAll(difficulty string) ([]structs.Question, error) {
	var questions []structs.Question

	query := s.db
	if difficulty != "" {
		query = query.Where("difficulty = ?", difficulty)
	}

	err := query.Find(&questions).Error
	return questions, err
}

func (s *QuestionService) GetByID(id int) (*structs.Question, error) {
	if id <= 0 {
		return nil, errors.New("invalid id")
	}

	var q structs.Question
	if err := s.db.First(&q, id).Error; err != nil {
		return nil, err
	}
	return &q, nil
}
