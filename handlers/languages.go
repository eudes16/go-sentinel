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
		applications.POST("/", insertLanguages)
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

func insertLanguages(c *gin.Context) {
	var lang entities.Language
	c.BindJSON(&lang)
	database.Connector.Create(&lang)
}
