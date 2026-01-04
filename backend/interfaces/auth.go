package interfaces

import s "techno/backend/structs"

type AuthService interface {
	Login(req LoginRequest) (*s.User, error)
}

type UserRepository interface {
	GetByEmail(email string) (*s.User, error)
	CreateUser(user *s.User) error
}

type GoogleClient interface {
	VerifyToken(token string) (*s.User, error)
}
