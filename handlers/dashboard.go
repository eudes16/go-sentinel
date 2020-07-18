package handlers

import (
	"eudes16/go-sentinel/database"
	"eudes16/go-sentinel/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DashboardHandlers(r *gin.Engine) {
	dashboard := r.Group("/dashboard")
	{
		dashboard.GET("/", getAllDashboards)
	}
}

func getAllDashboards(c *gin.Context) {
	var dashboard []entities.DashBoard

	database.Connector.Raw("SELECT COUNT(*), a.id, a.name FROM logs LEFT JOIN applications a ON a.id = logs.application_id GROUP BY a.id").Scan(&dashboard)

	c.JSON(http.StatusOK, dashboard)
}
