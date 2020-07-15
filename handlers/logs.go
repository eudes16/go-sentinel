package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// LogsHandlers ...
func LogsHandlers(r *gin.Engine) {
	logs := r.Group("/logs")
	{
		logs.GET("/", logsIndex)
		logs.POST("/", logsCreate)
	}
}

// Method GET
func logsIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Logs",
	})
}

// Method POST
func logsCreate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Logs POST",
	})
}
