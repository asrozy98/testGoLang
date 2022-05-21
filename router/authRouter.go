package router

import (
	"testGoLang/handler"
	"testGoLang/repository"
	"testGoLang/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRouter(route *gin.Engine, db *gorm.DB) {
	authRepository := repository.NewAuthRepo(db)
	authService := service.NewAuthService(authRepository)
	authHandler := handler.NewAuthHandler(authService)
	authRoter := route.Group("/api/auth")
	{
		authRoter.POST("/login", authHandler.LoginHandler)
		authRoter.POST("/register", authHandler.RegisterHandler)
	}
}
