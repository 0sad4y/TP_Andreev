package models

type Employee struct {
	ID            uint               `gorm:"primaryKey"`
	Name          string             `gorm:"type:text;not null"`
	Assignments   []AssignmentToTrip `gorm:"foreignKey:EmployeeID"`
	BusinessTrips []BusinessTrip     `gorm:"many2many:assignment_to_trips;joinForeignKey:EmployeeID;References:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
