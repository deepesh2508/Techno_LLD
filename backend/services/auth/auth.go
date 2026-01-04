package auth

import (
	"errors"
	i "techno/backend/interfaces"
	s "techno/backend/structs"
)

type EmailOTPAuthService struct {
	userRepo i.UserRepository
	otpSvc   OTPService
}

func (e *EmailOTPAuthService) Login(req i.LoginRequest) (*s.User, error) {
	if req.Email == "" || req.OTP == "" {
		return nil, errors.New("email and otp required")
	}

	// OTP verification
	if !e.otpSvc.Verify(req.Email, req.OTP) {
		return nil, errors.New("invalid or expired otp")
	}

	user, err := e.userRepo.GetByEmail(req.Email)
	if err == nil {
		return user, nil
	}

	// First time login → auto signup
	newUser := &s.User{
		Email: req.Email,
		// Provider: "email_otp",
		Status: "active",
	}

	if err := e.userRepo.CreateUser(newUser); err != nil {
		return nil, err
	}

	return newUser, nil
}

type GoogleAuthService struct {
	userRepo i.UserRepository
	google   i.GoogleClient
}

func (g *GoogleAuthService) Login(req i.LoginRequest) (*s.User, error) {
	if req.Token == "" {
		return nil, errors.New("google token required")
	}

	googleUser, err := g.google.VerifyToken(req.Token)
	if err != nil {
		return nil, err
	}

	user, err := g.userRepo.GetByEmail(googleUser.Email)
	if err == nil {
		return user, nil
	}

	newUser := &s.User{
		Email: googleUser.Email,
		Name:  googleUser.Name,
		// Provider: "google",
		Status: "active",
	}

	if err := g.userRepo.CreateUser(newUser); err != nil {
		return nil, err
	}

	return newUser, nil
}

type AuthStrategyFactory struct {
	emailOTP i.AuthService
	google   i.AuthService
}

func (f *AuthStrategyFactory) GetStrategy(provider string) i.AuthService {
	switch provider {
	case "email_otp":
		return f.emailOTP
	case "google":
		return f.google
	default:
		return nil
	}
}

func LoginHandler(factory *AuthStrategyFactory, provider string, req i.LoginRequest) (*s.User, error) {

	strategy := factory.GetStrategy(provider)
	if strategy == nil {
		return nil, errors.New("invalid login provider")
	}

	return strategy.Login(req)
}
