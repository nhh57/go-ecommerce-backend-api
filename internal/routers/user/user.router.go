package user

import (
	"github.com/gin-gonic/gin"
	"github.com/nhh57/go-ecommerce-backend-api/internal/controller/account"
)

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public router
	//this is non-dependency
	//ur := repo.NewUserRepository()
	//us := service.NewUserService(ur)
	//userHandlerNonDependency := controller.NewUserController(us)
	//userController, _ := wire.InitUserRouterHanlder()
	// Wire Go
	// Dependency injection
	userRouterPublic := Router.Group("/user")
	{
		//userRouterPublic.POST("/register", userController.Register) // register -> yes -> no
		userRouterPublic.POST("/register", account.Login.Register) // register -> yes -> no

		userRouterPublic.GET("/login", account.Login.Login) // login -> YES -> No
		userRouterPublic.GET("/otp")                        //
	}
	// private router
	userRouterPrivate := Router.Group("/user")
	//userRouterPrivate.Use(limiter())
	//userRouterPrivate.Use(Authen())
	//userRouterPrivate.Use(Permissions())
	{
		userRouterPrivate.POST("/get_info")
	}

}
