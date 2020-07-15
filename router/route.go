package router

import (
	"eudes16/go-sentinel/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Attach ...
func Attach(router *gin.Engine) {
	// Abilita o CORS
	router.Use(cors.Default())

	// declaração das rotas
	handlers.HomeHandlers(router)

	handlers.LogsHandlers(router)
}
