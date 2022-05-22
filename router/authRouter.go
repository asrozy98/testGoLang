package router

import (
	"testGoLang/handler"
	"testGoLang/middleware"
	"testGoLang/repository"
	"testGoLang/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRouter(route *gin.Engine, db *gorm.DB) {
	authRepository := repository.NewAuthRepo(db)
	authService := service.NewAuthService(authRepository)
	authHandler := handler.NewAuthHandler(authService)
	authRouter := route.Group("/api/auth")
	// authRouter.Use()
	{
		authRouter.POST("/login", authHandler.LoginHandler)
		authRouter.POST("/register", authHandler.RegisterHandler)
		authRouter.GET("/profile", middleware.AuthJwtCheck(), authHandler.ProfileHandler)
	}
}
