package router

import (
	"github.com/Andesson/marketplace-auth-service/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	v1 := router.Group("/auth-service/v1")
	{
		v1.POST("/auth/signup", handler.CreateUserHandler)
		v1.POST("/auth/login", handler.CreateUserHandler)
		v1.GET("/user", handler.CreateUserHandler)
	}
}
