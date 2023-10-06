package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID `gorm:"primary_key"`
	// ID     uint   `gorm:"primary_key"`
	Name      string `gorm:"type:varchar(255);not null"`
	Email     string `gorm:"uniqueIndex;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}