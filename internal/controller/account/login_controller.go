package account

import (
	"github.com/gin-gonic/gin"
	"github.com/nhh57/go-ecommerce-backend-api/global"
	"github.com/nhh57/go-ecommerce-backend-api/internal/model"
	"github.com/nhh57/go-ecommerce-backend-api/internal/service"
	"github.com/nhh57/go-ecommerce-backend-api/pkg/response"
	"go.uber.org/zap"
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

func (c *cUserLogin) Register(ctx *gin.Context) {
	var params model.RegisterInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}
	codeStatus, err := service.UserLogin().Register(ctx, &params)
	if err != nil {
		global.Logger.Error("Error registering user OTP", zap.Error(err))
		response.ErrorResponse(ctx, codeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, codeStatus, nil)
}
