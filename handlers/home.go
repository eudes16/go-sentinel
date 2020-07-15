package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HomeHandlers ...
func homeHandlers(r *gin.Engine) {
	home := r.Group("/")
	{
		home.GET("/", homeIndex)
	}
}

func homeIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Home",
	})
}