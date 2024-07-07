package initialize

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	c "github.com/nhh57/go-ecommerce-backend-api/internal/controller"
	"github.com/nhh57/go-ecommerce-backend-api/internal/middlewares"
)

func AA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --> AA")
		c.Next()
		fmt.Println("Alter --> AA")
	}
}

func BB() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --> BB")
		c.Next()
		fmt.Println("Alter --> BB")
	}
}

func CC(c *gin.Context) {
	fmt.Println("Before --> CC")
	c.Next()
	fmt.Println("Alter --> CC")
}

func InitRouter() *gin.Engine {
	r := gin.Default()
	// use the middleware
	r.Use(middlewares.AuthenMiddleware(), BB(), CC)
	v1 := r.Group("/v1/2024")
	{
		v1.GET("/ping", c.NewPongController().Pong)
		v1.GET("/user/1", c.NewUserController().GetUserByID)
		v1.PUT("/ping", Pong)
		v1.PATCH("/ping", Pong)
		v1.DELETE("/ping", Pong)
	}
	return r

}

func Pong(c *gin.Context) {
	name := c.Param("name")
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong" + name,
		"uid":     uid,
	})
}
