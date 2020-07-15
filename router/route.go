package router

import (
	"eudes16/go-sentinel/handlers"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
)

// Attach ...
func Attach(router *gin.Engine) {
	// Abilita o CORS
	router.Use(cors.Middleware(cors.Config{
		Origins:        "*",
	}))

	// declaração das rotas
	handlers.HomeHandlers(router)

	handlers.LogsHandlers(router)
}
