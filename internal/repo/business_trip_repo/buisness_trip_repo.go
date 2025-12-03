package business_trip_repo

import (
	"TP_Andreev/internal/dto"
	"TP_Andreev/internal/models"

	"gorm.io/gorm"
)

type BusinessTripRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) *BusinessTripRepo {
	return &BusinessTripRepo{db: db}
}

func (repo *BusinessTripRepo) All() (*[]dto.BuisnessTripDTO, error) {
	var businessTrips []models.BusinessTrip
	err := repo.db.Model(&models.BusinessTrip{}).Preload("Assignments").Preload("Assignments.Employee").Find(&businessTrips).Error

	var result []dto.BuisnessTripDTO

	for _, b := range businessTrips {
		businessTripDTO := dto.BuisnessTripDTO{
			ID:          b.ID,
			Destination: b.Destination,
			StartAt:     b.StartAt,
			EndAt:       b.EndAt,
		}

		var employeeTrips []dto.EmployeeTripDTO
		for _, a := range b.Assignments {
			employeeDTO := dto.EmployeeDTO{
				ID:   a.Employee.ID,
				Name: a.Employee.Name,
			}

			trip := dto.EmployeeTripDTO{
				MoneySpent: a.MoneySpent,
				Employee:   employeeDTO,
			}
			employeeTrips = append(employeeTrips, trip)
		}

		businessTripDTO.Employees = employeeTrips
		result = append(result, businessTripDTO)
	}

	return &result, err
}
