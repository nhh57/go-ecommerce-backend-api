package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/nhh57/go-ecommerce-backend-api/global"
	"github.com/nhh57/go-ecommerce-backend-api/internal/routers"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Model == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}
	// middlewares
	r.Use() //logger
	r.Use() // cross
	r.Use() // limiter global
	manageRouter := routers.RouterGroupApp.Manage
	userRouter := routers.RouterGroupApp.User

	MainGroup := r.Group("/v1/2024")
	{
		MainGroup.GET("/checkStatus") // tracking monitor
	}
	{
		userRouter.InitUserRouter(MainGroup)
		userRouter.InitProductRouter(MainGroup)
	}
	{
		manageRouter.InitUserRouter(MainGroup)
		manageRouter.InitAdminRouter(MainGroup)
	}
	return r
}
