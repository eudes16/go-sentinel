package entities

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

// Event is used by pop to map your .model.Name.Proper.Pluralize.Underscore models table to your go code.
type Event struct {
	ID           uuid.UUID `gorm:"type:uuid;PRIMARY_KEY;"`
	LogID        uuid.UUID `gorm:"type:uuid; foreign_key:id"`
	Log          Log
	EventTypeID  uuid.UUID `gorm:"type:uuid; foreign_key:id"`
	EventType    EventType
	ReceivedURL  string    `gorm:"type:varchar(255)"`
	SendedURL    string    `gorm:"type:varchar(255)"`
	Method       string    `gorm:"type:varchar(255)"`
	ResquestBody string    `gorm:"type:varchar(255)"`
	ResponseBody string    `gorm:"type:varchar(255)"`
	CreatedAt    time.Time `gorm:"type:timestamp"`
	UpdatedAt    time.Time `gorm:"type:timestamp"`
}
