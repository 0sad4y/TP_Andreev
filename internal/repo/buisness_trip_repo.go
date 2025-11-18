package repo

import (
	"TP_Andreev/internal/dto"
	"TP_Andreev/internal/models"
	"fmt"

	"gorm.io/gorm"
)

type BusinessTripRepo struct {
	DB *gorm.DB
}

func (r *BusinessTripRepo) First() error {
	
	return nil
}

func (r *BusinessTripRepo) All() ([]dto.BuisnessTripDTO, error) {
	var assignments []models.AssignmentToTrip
	err := r.DB.Model(&models.AssignmentToTrip{}).Preload("Employee").Preload("BusinessTrip").Find(&assignments).Error
	fmt.Print()

	return nil, err
}
