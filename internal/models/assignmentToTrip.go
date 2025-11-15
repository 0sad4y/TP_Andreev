package models

import (
	"gorm.io/gorm"
)

type AssignmentToTrip struct {
	gorm.Model
	ID           uint `gorm:"primaryKey"`
	MoneySpend   int
	Employee     Employee     `gorm:"foreignKey"`
	BuisnessTrip BuisnessTrip `gorm:"foreignKey"`
}
