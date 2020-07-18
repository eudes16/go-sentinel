package main

import (
	"encoding/json"
	"eudes16/go-sentinel/database"
	"eudes16/go-sentinel/entities"
	"eudes16/go-sentinel/router"
	"eudes16/go-sentinel/utils"
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

	utils.InitBot()

	r.Run(":3000")
}

func runMigrations() {
	database.Connector.AutoMigrate(&entities.Language{}, &entities.Log{}, &entities.Application{}, entities.Event{}, &entities.EventType{})

	database.Connector.Model(&entities.Log{}).AddForeignKey("language_id", "languages(id)", "RESTRICT", "RESTRICT")
	database.Connector.Model(&entities.Log{}).AddForeignKey("application_id", "applications(id)", "RESTRICT", "RESTRICT")
	database.Connector.Model(&entities.Event{}).AddForeignKey("log_id", "logs(id)", "RESTRICT", "RESTRICT")
	database.Connector.Model(&entities.Event{}).AddForeignKey("event_type_id", "event_types(id)", "RESTRICT", "RESTRICT")


	type JsonLanguage struct {
		Language []entities.Language `json:"languages"`
	}

	bytes := utils.LoadLanguages("assets/languages.json")

	var jsonLanguage JsonLanguage

	if err := json.Unmarshal(bytes, &jsonLanguage); err != nil {
		panic(err)
	}

	for _, l := range jsonLanguage.Language {
		database.Connector.Create(&l)
	}

	database.Connector.Model(&entities.Application{}).Exec("INSERT INTO applications (id, name, created_at, updated_at) VALUES (?, ?, ?, ?);", "0100ab8b-456b-4ef4-9298-aaa8ba5a87d6", "Vista CRM", time.Now(), time.Now())
	database.Connector.Model(&entities.Application{}).Exec("INSERT INTO applications (id, name, created_at, updated_at) VALUES (?, ?, ?, ?);", "aef14341-a7df-484d-b452-53139dcb2158", "Vista Mobile", time.Now(), time.Now())
	database.Connector.Model(&entities.Application{}).Exec("INSERT INTO applications (id, name, created_at, updated_at) VALUES (?, ?, ?, ?);", "173e815c-1ec9-4da5-a54c-f5d2825507cd", "Vista Office", time.Now(), time.Now())
	database.Connector.Model(&entities.Application{}).Exec("INSERT INTO applications (id, name, created_at, updated_at) VALUES (?, ?, ?, ?);", "a3e93130-bd9b-45e1-b43d-24a17cac2766", "Vista Analytics", time.Now(), time.Now())

	database.Connector.Model(&entities.Application{}).Exec("INSERT INTO event_types (id, name, created_at, updated_at) VALUES (?, ?, ?, ?);", uuid.NewV4(), "Catch Error", time.Now(), time.Now())
	database.Connector.Model(&entities.Application{}).Exec("INSERT INTO event_types (id, name, created_at, updated_at) VALUES (?, ?, ?, ?);", uuid.NewV4(), "Send Mail", time.Now(), time.Now())
	database.Connector.Model(&entities.Application{}).Exec("INSERT INTO event_types (id, name, created_at, updated_at) VALUES (?, ?, ?, ?);", uuid.NewV4(), "Speak on Discord", time.Now(), time.Now())
	database.Connector.Model(&entities.Application{}).Exec("INSERT INTO event_types (id, name, created_at, updated_at) VALUES (?, ?, ?, ?);", uuid.NewV4(), "Speak on Rocket", time.Now(), time.Now())

}
