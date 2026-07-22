package service

import (
	"context"

	"github.com/1001001010/messanger/gen/auth"
	"github.com/1001001010/messanger/internal/database"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	db *database.Queries
}

func (s *AuthService) Register(
	ctx context.Context,
	req *auth.RegisterRequest,
) (*auth.RegisterResponse, error) {

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return nil, err
	}

	user, err := s.db.CreateUser(ctx, database.CreateUserParams{
		Email:        req.Email,
		FirstName:    req.FirstName,
		PasswordHash: string(hash),
	})

	if err != nil {
		return nil, err
	}

	return &auth.RegisterResponse{
		User: mapUser(user),
	}, nil
}
