package auth

import (
	"context"
	"errors"
	s "techno/backend/structs"

	"google.golang.org/api/idtoken"
)

type GoogleOAuthClient struct {
	audience string // Google OAuth client ID
}

func NewGoogleOAuthClient(clientID string) *GoogleOAuthClient {
	return &GoogleOAuthClient{audience: clientID}
}

func (g *GoogleOAuthClient) VerifyToken(token string) (*s.User, error) {
	payload, err := idtoken.Validate(context.Background(), token, g.audience)
	if err != nil {
		return nil, errors.New("invalid google token")
	}

	email, _ := payload.Claims["email"].(string)
	name, _ := payload.Claims["name"].(string)

	if email == "" {
		return nil, errors.New("email not present in token")
	}

	return &s.User{
		Email: email,
		Name:  name,
	}, nil
}
