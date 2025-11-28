package repo

import (
	"TP_Andreev/internal/dto"
	"TP_Andreev/internal/models"

	"gorm.io/gorm"
)

type EmployeeRepo struct {
	DB *gorm.DB
}

func (repo *EmployeeRepo) Find(id uint) (*dto.EmployeeDTO, error) {
	var employee models.Employee
	err := repo.DB.Model(&models.Employee{}).Preload("Assignments").Preload("Assignments.BusinessTrip").Find(&employee, id).Error

	employeeDTO := dto.EmployeeDTO{
		ID:   employee.ID,
		Name: employee.Name,
	}

	var employeeTrips []dto.EmployeeTripDTO
	for _, a := range employee.Assignments {
		businessTripDTO := dto.BuisnessTripDTO{
			ID:          a.BusinessTrip.ID,
			Destination: a.BusinessTrip.Destination,
			StartAt:     a.BusinessTrip.StartAt,
			EndAt:       a.BusinessTrip.EndAt,
		}
		trip := dto.EmployeeTripDTO{
			MoneySpent:   a.MoneySpent,
			Employee:     employeeDTO,
			BuisnessTrip: businessTripDTO,
		}
		employeeTrips = append(employeeTrips, trip)
	}

	employeeDTO.Trips = employeeTrips

	return &employeeDTO, err
}

func (repo *EmployeeRepo) All() (*[]dto.EmployeeDTO, error) {
	var employees []models.Employee
	err := repo.DB.Model(&models.Employee{}).Preload("Assignments").Preload("Assignments.BusinessTrip").Find(&employees).Error

	var result []dto.EmployeeDTO

	for _, e := range employees {
		employeeDTO := dto.EmployeeDTO{
			ID:   e.ID,
			Name: e.Name,
		}

		var employeeTrips []dto.EmployeeTripDTO
		for _, a := range e.Assignments {
			businessTripDTO := dto.BuisnessTripDTO{
				ID:          a.BusinessTrip.ID,
				Destination: a.BusinessTrip.Destination,
				StartAt:     a.BusinessTrip.StartAt,
				EndAt:       a.BusinessTrip.EndAt,
			}
			trip := dto.EmployeeTripDTO{
				MoneySpent:   a.MoneySpent,
				Employee:     employeeDTO,
				BuisnessTrip: businessTripDTO,
			}
			employeeTrips = append(employeeTrips, trip)
		}

		employeeDTO.Trips = employeeTrips
		result = append(result, employeeDTO)
	}

	return &result, err
}
