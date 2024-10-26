package initialize

import (
	"github.com/nhh57/go-ecommerce-backend-api/global"
	"github.com/nhh57/go-ecommerce-backend-api/internal/database"
	"github.com/nhh57/go-ecommerce-backend-api/internal/service"
	"github.com/nhh57/go-ecommerce-backend-api/internal/service/impl"
)

func InitServiceInterface() {
	queries := database.New(global.Mdbc)
	// User Service Interface
	service.InitUserLogin(impl.NewUserLoginImpl(queries))
	//................................................................
}
