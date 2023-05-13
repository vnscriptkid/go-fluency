package models

import (
	"github.com/google/uuid"
)

type Range struct {
	ID      uuid.UUID `gorm:"primary_key;type:uuid"`
	Min     int64     `gorm:"not null"`
	Max     int64     `gorm:"not null"`
	Current int64     `gorm:"not null"`
}
