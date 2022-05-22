package handler

import (
	"fmt"
	"net/http"
	"testGoLang/model"
	"testGoLang/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type authHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *authHandler {
	return &authHandler{authService}
}

func (handler *authHandler) LoginHandler(c *gin.Context) {
	var loginRequest model.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		errorMessages := []any{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on %s, because: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Validation error",
			"error":   errorMessages,
		})
		return
	}

	result, err := handler.authService.Login(loginRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": result,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"token":   result,
			"expires": "10 minutes",
		},
	})

}

func (handler *authHandler) RegisterHandler(c *gin.Context) {
	var registerRequest model.RegisterRequest

	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		errorMessages := []any{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on %s, because: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Validation error",
			"error":   errorMessages,
		})
		return
	}

	result, err := handler.authService.Register(registerRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": result,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Register success",
		"data":    result,
	})
}

func (handler *authHandler) ProfileHandler(c *gin.Context) {
	email := c.MustGet("email").(string)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    gin.H{"email": email},
	})
}
