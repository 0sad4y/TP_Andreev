package repo

import (
	"TP_Andreev/internal/models"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type EmployeeRepo struct {
	DB *gorm.DB
}

func (e *EmployeeRepo) First() error {
	var employee models.Employee
	err := e.DB.Preload("BusinessTrips").Preload("BusinessTrips.Assignments").First(&employee).Error;
	fmt.Print(employee)

	return err
}

func (e *EmployeeRepo) All() {
	var employees []models.Employee
	if err := e.DB.Find(&employees).Error; err != nil {
		log.Fatalf("Failed to fetch all employees: %v", err)
	}
}