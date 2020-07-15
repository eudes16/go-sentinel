package main

import (
	"eudes16/go-sentinel/database"
	"eudes16/go-sentinel/entities"
	"eudes16/go-sentinel/router"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"time"
)

func main() {

	gin.ForceConsoleColor()
	r := gin.Default()

	database.OpenConnection()
	//runMigrations()

	router.Attach(r)

	r.Run(":3000")
}

func runMigrations() {
	database.Connector.AutoMigrate(&entities.Language{}, &entities.Log{}, &entities.Application{}, entities.Event{}, &entities.EventType{})

	database.Connector.Model(&entities.Log{}).AddForeignKey("language_id", "languages(id)", "RESTRICT", "RESTRICT")
	database.Connector.Model(&entities.Log{}).AddForeignKey("application_id", "applications(id)", "RESTRICT", "RESTRICT")
	database.Connector.Model(&entities.Event{}).AddForeignKey("log_id", "logs(id)", "RESTRICT", "RESTRICT")
	database.Connector.Model(&entities.Event{}).AddForeignKey("event_type_id", "event_types(id)", "RESTRICT", "RESTRICT")

	database.Connector.Model(&entities.Application{}).Exec("INSERT INTO languages (id, name, created_at, updated_at) VALUES (?, ?, ?, ?);", uuid.NewV4(), "PHP", time.Now(), time.Now())
	database.Connector.Model(&entities.Application{}).Exec("INSERT INTO languages (id, name, created_at, updated_at) VALUES (?, ?, ?, ?);", uuid.NewV4(), "Dart", time.Now(), time.Now())
	database.Connector.Model(&entities.Application{}).Exec("INSERT INTO languages (id, name, created_at, updated_at) VALUES (?, ?, ?, ?);", uuid.NewV4(), "Java Script", time.Now(), time.Now())
	database.Connector.Model(&entities.Application{}).Exec("INSERT INTO languages (id, name, created_at, updated_at) VALUES (?, ?, ?, ?);", uuid.NewV4(), "Golang", time.Now(), time.Now())

	database.Connector.Model(&entities.Application{}).Exec("INSERT INTO applications (id, name, created_at, updated_at) VALUES (?, ?, ?, ?);", uuid.NewV4(), "Vista CRM", time.Now(), time.Now())
	database.Connector.Model(&entities.Application{}).Exec("INSERT INTO applications (id, name, created_at, updated_at) VALUES (?, ?, ?, ?);", uuid.NewV4(), "Vista Mobile", time.Now(), time.Now())
	database.Connector.Model(&entities.Application{}).Exec("INSERT INTO applications (id, name, created_at, updated_at) VALUES (?, ?, ?, ?);", uuid.NewV4(), "Vista Office", time.Now(), time.Now())
	database.Connector.Model(&entities.Application{}).Exec("INSERT INTO applications (id, name, created_at, updated_at) VALUES (?, ?, ?, ?);", uuid.NewV4(), "Vista Analytics", time.Now(), time.Now())

	database.Connector.Model(&entities.Application{}).Exec("INSERT INTO event_types (id, name, created_at, updated_at) VALUES (?, ?, ?, ?);", uuid.NewV4(), "Catch Error", time.Now(), time.Now())
	database.Connector.Model(&entities.Application{}).Exec("INSERT INTO event_types (id, name, created_at, updated_at) VALUES (?, ?, ?, ?);", uuid.NewV4(), "Send Mail", time.Now(), time.Now())
	database.Connector.Model(&entities.Application{}).Exec("INSERT INTO event_types (id, name, created_at, updated_at) VALUES (?, ?, ?, ?);", uuid.NewV4(), "Speak on Discord", time.Now(), time.Now())
	database.Connector.Model(&entities.Application{}).Exec("INSERT INTO event_types (id, name, created_at, updated_at) VALUES (?, ?, ?, ?);", uuid.NewV4(), "Speak on Rocket", time.Now(), time.Now())

}
