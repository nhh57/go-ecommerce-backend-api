package account

import (
	"github.com/gin-gonic/gin"
	"github.com/nhh57/go-ecommerce-backend-api/internal/service"
	"github.com/nhh57/go-ecommerce-backend-api/pkg/response"
)

// management controller Login User
var Login = new(cUserLogin)

type cUserLogin struct{}

func (c *cUserLogin) Login(ctx *gin.Context) {
	// Implement logic for login
	err := service.UserLogin().Login(ctx)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.ErrCodeSuccess, nil)
}
