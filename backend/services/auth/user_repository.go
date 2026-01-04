package auth

import (
	s "techno/backend/structs"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func (r userRepository) GetByEmail(email string) (*s.User, error) {
	var user s.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r userRepository) CreateUser(user *s.User) error {
	return r.db.Create(user).Error
}
