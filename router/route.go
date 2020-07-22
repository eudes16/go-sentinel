package router

import (
	"eudes16/go-sentinel/handlers"
	"github.com/gin-gonic/gin"
)

// Attach ...
func Attach(router *gin.Engine) {
	// declaração das rotas
	handlers.HomeHandlers(router)

	handlers.LogsHandlers(router)

	handlers.ApplicationsHandlers(router)

	handlers.DashboardHandlers(router)

	handlers.LanguageHandlers(router)
}
