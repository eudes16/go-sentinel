package entities

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

// Language ...
type Language struct {
	ID        uuid.UUID  `gorm:"type:uuid;PRIMARY_KEY;" json:"id"`
	Name      string     `gorm:"type:varchar(255)" json:"name"`
	CreatedAt time.Time  `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt time.Time  `gorm:"type:timestamp" json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}
