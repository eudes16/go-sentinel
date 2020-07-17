package handlers

import (
	"eudes16/go-sentinel/database"
	"eudes16/go-sentinel/entities"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ApplicationsHandlers(r *gin.Engine) {
	applications := r.Group("/applications")
	{
		applications.GET("/", applicationsGetAll)
		applications.GET("/logs/:application", logsByApplication)
	}
}

func applicationsGetAll(c *gin.Context) {
	var applications []entities.Application
	if err := database.Connector.Find(&applications).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, applications)
	}
}

func logsByApplication(c *gin.Context) {
	id, _ := c.Params.Get("application")
	var logs []entities.Log
	err := database.Connector.Where("application_id = ?", id).Find(&logs).Error

	for i, _ := range logs {
		database.Connector.Model(&logs[i]).Related(&logs[i].Application).Related(&logs[i].Language)
	}

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, logs)
}
