package service

import (
	"github.com/nhh57/go-ecommerce-backend-api/internal/repo"
	"github.com/nhh57/go-ecommerce-backend-api/pkg/response"
)

type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo repo.IUserRepository
}

func NewUserService(
	userRepo repo.IUserRepository,
) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (us *userService) Register(email string, purpose string) int {
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrCodeUserHasExists
	}
	return response.ErrCodeSuccess
}
