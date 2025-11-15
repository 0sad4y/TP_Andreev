package models

import (
	"time"

	"gorm.io/gorm"
)

type BuisnessTrip struct {
	gorm.Model
	ID          uint `gorm:"primaryKey"`
	Destination string
	StartAt     time.Time
	EndAt       time.Time
}
