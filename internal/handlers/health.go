package handlers

import "github.com/gin-gonic/gin"

func HealthHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Hng stage 1 task",
	})
}
