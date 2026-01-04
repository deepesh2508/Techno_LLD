package interfaces

import s "techno/backend/structs"

type QuestionService interface {
	Create(q *s.Question) error
	GetAll(difficulty string) ([]s.Question, error)
	GetByID(id int) (*s.Question, error)
}
