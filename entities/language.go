package entities

import (
	"github.com/google/uuid"
	"time"
)

// Language ...
type Language struct {
	ID        uuid.UUID  `gorm:"type:uuid;PRIMARY_KEY;"`
	Name      string     `gorm:"type:varchar(255)"`
	CreatedAt time.Time  `gorm:"type:timestamp"`
	UpdatedAt time.Time  `gorm:"type:timestamp"`
	DeletedAt *time.Time `sql:"index"`
}
