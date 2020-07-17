package entities

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"time"
)

// Log is used by pop to map your logs models table to your go code.
type Log struct {
	ID            uuid.UUID   `gorm:"type:uuid;PRIMARY_KEY;" json:"id"`
	Code          string      `gorm:"type:varchar(255)" json:"code"`
	Message       string      `gorm:"type:text" json:"message"`
	StackTrace    string      `gorm:"type:text" json:"stacktrace"`
	Language      Language    `json:"language"`
	LanguageID    uuid.UUID   `gorm:"type:uuid;" json:"language_id,omitempty"`
	Application   Application `json:"application"`
	ApplicationID uuid.UUID   `gorm:"type:uuid;" json:"application_id,omitempty"`
	File          string      `gorm:"type:text" json:"file"`
	UserID        string      `gorm:"type:varchar(255)" json:"user_id"`
	ClientID      string      `gorm:"type:varchar(255)" json:"client_id"`
	CreatedAt     time.Time   `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt     time.Time   `gorm:"type:timestamp" json:"updated_at"`
	DeletedAt     *time.Time  `sql:"index" json:"deleted_at"`
}

func (base *Log) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("ID", uuid)
}
