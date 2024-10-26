package impl

import (
	"context"
	"github.com/nhh57/go-ecommerce-backend-api/internal/database"
)

type sUserLogin struct {
	// Implement the IUserLogin interface here.
	r *database.Queries
}

func NewUserLoginImpl(r *database.Queries) *sUserLogin {
	return &sUserLogin{
		r: r,
	}
}

func (s *sUserLogin) Login(ctx context.Context) error {
	// Implement login logic here.
	return nil
}

func (s *sUserLogin) Register(ctx context.Context) error {
	// Implement login logic here.
	return nil
}

func (s *sUserLogin) VerifyOTP(ctx context.Context) error {
	// Implement login logic here.
	return nil
}

func (s *sUserLogin) UpdatePasswordRegister(ctx context.Context) error {
	// Implement login logic here.
	return nil
}
