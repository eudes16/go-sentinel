package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Router ...
func Router(router *gin.Engine) {
	// Abilita o CORS
	router.Use(cors.Default())

	// declaração das rotas
	handlers.homeHandlers(router)

	handlers.logsHandlers(router)
}
