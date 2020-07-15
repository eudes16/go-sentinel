package entities

import (
	"github.com/google/uuid"
)

// Language ...
type Language struct {
	ID   uuid.UUID `gorm:"type:uuid;PRIMARY_KEY;"`
	Name string    `gorm:"type:varchar(255)"`
}
