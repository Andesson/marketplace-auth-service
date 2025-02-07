package router

import (
	"github.com/Andesson/marketplace-auth-service/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	v1 := router.Group("/auth-service/v1")
	{
		v1.POST("/login", handler.CreateUserHandler)
	}
}
