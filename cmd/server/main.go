package main

import (
	_ "github.com/nhh57/go-ecommerce-backend-api/cmd/swag/docs"
	"github.com/nhh57/go-ecommerce-backend-api/internal/initialize"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title           API Documentation Ecommerce Backend SHOPDEVGO
// @version         1.0.0
// @description     This is a sample server caller server.
// @termsOfService	github.com/nhh57/go-ecommerce-backend-api

// @contact.name   TEAM TIPSGO
// @contact.url    github.com/nhh57/go-ecommerce-backend-api
// @contact.email  tipsgo@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8082
// @BasePath  /v1/2024
// @schema http
func main() {
	r := initialize.Run()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8082")
}
