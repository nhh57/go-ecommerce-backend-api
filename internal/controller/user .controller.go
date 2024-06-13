package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/nhh57/go-ecommerce-backend-api/internal/service"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {

	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetUserByID(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"mesage": uc.userService.GetInfoUser(),
		"users":  []string{"ck7", "m10", "hainh"},
	})
}
