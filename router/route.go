package router

import (
	"eudes16/go-sentinel/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Attach ...
func Attach(router *gin.Engine) {
	// Abilita o CORS
	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token", "access-control-allow-origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		//AllowOrigins: []string{"*"},
		AllowAllOrigins: true,
		//AllowOriginFunc:  func(origin string) bool { return true },
		MaxAge:           86400,
	}))

	// declaração das rotas
	handlers.HomeHandlers(router)

	handlers.LogsHandlers(router)

	handlers.ApplicationsHandlers(router)

	handlers.DashboardHandlers(router)

	handlers.LanguageHandlers(router)
}
