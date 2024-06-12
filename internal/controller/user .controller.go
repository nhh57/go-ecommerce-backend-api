package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
type UserController struct {}

func NewUserController() *UserController {

return &UserController{}
}

func (u*UserController)GetUserByID(c *gin.Context) {
	name := c.DefaultQuery("name", "hainh")
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"mesage": "pong.hhhh..ping" + name,
		"uid":    uid,
		"users":  []string{"ck7","m10","hainh"},
	})
}