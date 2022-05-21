package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(db *gorm.DB) {
	route := gin.Default()

	route.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello word",
		})
	})

	AuthRouter(route, db)

	route.Run(":8000")
}
