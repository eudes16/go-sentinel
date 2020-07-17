package handlers

import (
	"eudes16/go-sentinel/database"
	"eudes16/go-sentinel/entities"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LogsHandlers ...
func LogsHandlers(r *gin.Engine) {
	logs := r.Group("/logs")
	{
		logs.GET("/", logsGetAll)
		logs.GET("/:id", logsById)
		logs.POST("/", logsCreate)
		logs.PUT("/:id", logsUpdate)
		logs.DELETE("/:id", logsDelete)
	}
}

// Method GET
func logsGetAll(c *gin.Context) {
	var logs []entities.Log
	if err := database.Connector.Find(&logs).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		for i, _ := range logs {
			database.Connector.Model(&logs[i]).Related(&logs[i].Application).Related(&logs[i].Language)
		}

		c.JSON(http.StatusOK, logs)
	}
}

func logsById(c *gin.Context) {
	id, _ := c.Params.Get("id")
	var log entities.Log
	err := database.Connector.Where("id = ?", id).Find(&log).Error

	database.Connector.Model(&log).Related(&log.Application).Related(&log.Language)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, log)
}

func logsUpdate(c *gin.Context) {
	var log entities.Log
	id := c.Params.ByName("id")
	if err := database.Connector.Where("id = ?", id).First(&log).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&log)
	database.Connector.Save(&log)

	database.Connector.Model(&log).Related(&log.Application).Related(&log.Language)

	c.JSON(200, log)
}

// Method POST
func logsCreate(c *gin.Context) {
	var log entities.Log
	c.BindJSON(&log)

	database.Connector.Create(&log)

	err := database.Connector.Model(&log).Related(&log.Application).Related(&log.Language).Error

	if err != nil {
		c.JSON(http.StatusNotModified, map[string]string{"message": "Not created"})
		return
	}

	//utils.SendMessage(&log)

	c.JSON(http.StatusCreated, log)
}

func logsDelete(c *gin.Context) {

	id, _ := c.Params.Get("id")
	var log entities.Log
	database.Connector.Where("id = ?", id).Delete(&log)
	data := fmt.Sprintf("id %v/", id)
	c.JSON(200, gin.H{data: "deleted"})
}
