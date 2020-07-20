package handlers

import (
	"eudes16/go-sentinel/database"
	"eudes16/go-sentinel/entities"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LanguageHandlers(r *gin.Engine) {
	applications := r.Group("/languages")
	{
		applications.GET("/", languagesGetAll)
	}
}

func languagesGetAll(c *gin.Context) {
	var languages []entities.Language
	if err := database.Connector.Find(&languages).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, languages)
	}
}