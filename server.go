package main

import (
	"fmt"

	"eudes16/go-sentinel/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var err error

func main() {
	// db.Connector, err = gorm.Open("postgres", db.ConnectorURL(db.BuildConnectorConfig()))

	if err != nil {
		fmt.Println("Status:", err)
	}

	// defer db.Connector.Close()
	// db.Connector.AutoMigrate(&entities.Language{})

	gin.ForceConsoleColor()
	r := gin.Default()

	r.Use(cors.Default())

	router.Attach(r)

	r.Run(":3000")
}
