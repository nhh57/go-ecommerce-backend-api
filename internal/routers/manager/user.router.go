package manager

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public router

	// private router
	userRouterPrivate := Router.Group("/admin")
	//userRouterPrivate.Use(limiter())
	//userRouterPrivate.Use(Authen())
	//userRouterPrivate.Use(Permissions())
	{
		userRouterPrivate.POST("/active_user")
	}

}
