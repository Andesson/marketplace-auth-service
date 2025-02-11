package router

import (
	"github.com/Andesson/marketplace-auth-service/handler"
	"github.com/Andesson/marketplace-auth-service/middleware"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	v1 := router.Group("/api")
	{
		v1.POST("/auth/signup", handler.CreateUserHandler)
		v1.POST("/auth/logon", handler.Logon)
	}

	// Grupo de rotas protegidas
	protected := router.Group("/auth-service/v1")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/protected", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "Você acessou uma rota protegida!"})
		})
	}

}
