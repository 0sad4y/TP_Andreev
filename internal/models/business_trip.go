package models

import "time"

type BusinessTrip struct {
	ID          uint               `gorm:"primaryKey"`
	Destination string             `gorm:"type:text;not null"`
	StartAt     time.Time          `gorm:"type:date;not null"`
	EndAt       time.Time          `gorm:"type:date;not null"`
	Assignments []AssignmentToTrip `gorm:"foreignKey:BusinessTripID"`
	// Employees   []Employee         `gorm:"many2many:assignment_to_trips;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Employees   []Employee         `gorm:"many2many:assignment_to_trip;joinForeignKey:BusinessTripID;References:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
