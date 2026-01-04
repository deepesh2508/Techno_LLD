package attempt

import (
	s "techno/backend/structs"

	"gorm.io/gorm"
)

type AttemptRepoDB struct {
	db *gorm.DB
}

func NewAttemptRepoDB(db *gorm.DB) *AttemptRepoDB {
	return &AttemptRepoDB{db: db}
}

func (r *AttemptRepoDB) Create(attempt *s.Attempt) error {
	return r.db.Create(attempt).Error
}

func (r *AttemptRepoDB) FindByID(id string) (*s.Attempt, error) {
	var attempt s.Attempt
	if err := r.db.First(&attempt, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &attempt, nil
}

func (r *AttemptRepoDB) Update(attempt *s.Attempt) error {
	return r.db.Save(attempt).Error
}
