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

	//utils.InitBot()

	r.Run(":3000")
}

func runMigrations() {
	database.Connector.AutoMigrate(&entities.Language{}, &entities.Log{}, &entities.Application{}, entities.Event{}, &entities.EventType{})

	database.Connector.Model(&entities.Log{}).AddForeignKey("language_id", "languages(id)", "RESTRICT", "RESTRICT")
	database.Connector.Model(&entities.Log{}).AddForeignKey("application_id", "applications(id)", "RESTRICT", "RESTRICT")
	database.Connector.Model(&entities.Event{}).AddForeignKey("log_id", "logs(id)", "RESTRICT", "RESTRICT")
	database.Connector.Model(&entities.Event{}).AddForeignKey("event_type_id", "event_types(id)", "RESTRICT", "RESTRICT")

	database.Connector.Model(&entities.Application{}).Exec("INSERT INTO languages (id, name, created_at, updated_at) VALUES (?, ?, ?, ?);", "052e1762-26b3-479a-b609-62efaf08cd21", "PHP", time.Now(), time.Now())
	database.Connector.Model(&entities.Application{}).Exec("INSERT INTO languages (id, name, created_at, updated_at) VALUES (?, ?, ?, ?);", "f4993d6d-1597-42cc-affa-58fbd0a83c58", "Dart", time.Now(), time.Now())
	database.Connector.Model(&entities.Application{}).Exec("INSERT INTO languages (id, name, created_at, updated_at) VALUES (?, ?, ?, ?);", "52000bc1-a701-40dc-8d6f-3c83157e5cd9", "Java Script", time.Now(), time.Now())
	database.Connector.Model(&entities.Application{}).Exec("INSERT INTO languages (id, name, created_at, updated_at) VALUES (?, ?, ?, ?);", "57f1cf80-a183-48d9-8702-f7c8f42083c5", "Golang", time.Now(), time.Now())

	database.Connector.Model(&entities.Application{}).Exec("INSERT INTO applications (id, name, created_at, updated_at) VALUES (?, ?, ?, ?);", "0100ab8b-456b-4ef4-9298-aaa8ba5a87d6", "Vista CRM", time.Now(), time.Now())
	database.Connector.Model(&entities.Application{}).Exec("INSERT INTO applications (id, name, created_at, updated_at) VALUES (?, ?, ?, ?);", "aef14341-a7df-484d-b452-53139dcb2158", "Vista Mobile", time.Now(), time.Now())
	database.Connector.Model(&entities.Application{}).Exec("INSERT INTO applications (id, name, created_at, updated_at) VALUES (?, ?, ?, ?);", "173e815c-1ec9-4da5-a54c-f5d2825507cd", "Vista Office", time.Now(), time.Now())
	database.Connector.Model(&entities.Application{}).Exec("INSERT INTO applications (id, name, created_at, updated_at) VALUES (?, ?, ?, ?);", "a3e93130-bd9b-45e1-b43d-24a17cac2766", "Vista Analytics", time.Now(), time.Now())

	database.Connector.Model(&entities.Application{}).Exec("INSERT INTO event_types (id, name, created_at, updated_at) VALUES (?, ?, ?, ?);", uuid.NewV4(), "Catch Error", time.Now(), time.Now())
	database.Connector.Model(&entities.Application{}).Exec("INSERT INTO event_types (id, name, created_at, updated_at) VALUES (?, ?, ?, ?);", uuid.NewV4(), "Send Mail", time.Now(), time.Now())
	database.Connector.Model(&entities.Application{}).Exec("INSERT INTO event_types (id, name, created_at, updated_at) VALUES (?, ?, ?, ?);", uuid.NewV4(), "Speak on Discord", time.Now(), time.Now())
	database.Connector.Model(&entities.Application{}).Exec("INSERT INTO event_types (id, name, created_at, updated_at) VALUES (?, ?, ?, ?);", uuid.NewV4(), "Speak on Rocket", time.Now(), time.Now())

}
