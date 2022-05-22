package middleware

import (
	"net/http"
	"strings"
	"testGoLang/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthJwtCheck() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Bad Request",
				"error":   "Request header auth not found",
			})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Bad Request",
				"error":   "Request header auth Incorrect format",
			})
			c.Abort()
			return
		}

		token := parts[1]
		claims := &model.Claims{}

		tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (any, error) {
			return model.JwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.JSON(http.StatusUnauthorized, gin.H{
					"success": false,
					"message": "Unauthorized",
					"error":   err.Error(),
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Bad Request",
				"error":   err.Error(),
			})
			c.Abort()
			return
		}
		if !tkn.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Token not valid",
				"error":   err.Error(),
			})
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		c.Next()
	}
}
