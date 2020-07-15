package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func logsHandlers(r *gin.Engine) {
	logs := r.Group("/logs")
	{
		logs.GET("/", logsIndex)
		logs.POST("/", logsIndex)
	}
}

func logsIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Logs",
	})
}

func logsCreate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Logs",
	})
}
