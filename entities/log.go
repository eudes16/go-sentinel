package entities

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"time"
)

// Log is used by pop to map your logs models table to your go code.
type Log struct {
	ID            uuid.UUID `gorm:"type:uuid;PRIMARY_KEY;"`
	Code          string    `gorm:"type:varchar(255)"`
	Message       string    `gorm:"type:text"`
	StackTrace    string    `gorm:"type:text"`
	Language      Language
	LanguageID    uuid.UUID `gorm:"type:uuid;"`
	Application   Application
	ApplicationID uuid.UUID  `gorm:"type:uuid;"`
	File          string     `gorm:"type:text"`
	UserID        string     `gorm:"type:varchar(255)"`
	ClientID      string     `gorm:"type:varchar(255)"`
	CreatedAt     time.Time  `gorm:"type:timestamp"`
	UpdatedAt     time.Time  `gorm:"type:timestamp"`
	DeletedAt     *time.Time `sql:"index"`
}

func (base *Log) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("ID", uuid)
}
