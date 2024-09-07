//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/nhh57/go-ecommerce-backend-api/internal/controller"
	"github.com/nhh57/go-ecommerce-backend-api/internal/repo"
	"github.com/nhh57/go-ecommerce-backend-api/internal/service"
)

func InitUserRouterHanlder() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepository,
		service.NewUserService,
		controller.NewUserController,
	)
	return new(controller.UserController), nil
}
