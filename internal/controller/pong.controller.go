package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongController struct{}

func NewPongController() *PongController {

	return &PongController{}
}

func (p *PongController) Pong(c *gin.Context) {
	name := c.DefaultQuery("name", "hainh")
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"mesage": "pong.hhhh..ping" + name,
		"uid":    uid,
		"users":  []string{"ck7", "m10", "hainh"},
	})
}
