package entities

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

// EventType is used by pop to map your .model.Name.Proper.Pluralize.Underscore models table to your go code.
type EventType struct {
	ID        uuid.UUID  `gorm:"type:uuid;PRIMARY_KEY;"`
	Name      string     `gorm:"type:varchar(255)"`
	CreatedAt time.Time  `gorm:"type:timestamp"`
	UpdatedAt time.Time  `gorm:"type:timestamp"`
	DeletedAt *time.Time `sql:"index"`
}
