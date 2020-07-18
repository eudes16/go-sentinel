package entities

import uuid "github.com/satori/go.uuid"

type DashBoard struct {
	Count       int
	ID          uuid.UUID
	Name        string
}
