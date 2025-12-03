package service

import (
	"TP_Andreev/internal/dto"
)

// AggregationStrategy defines how to extract and aggregate data from trips
type AggregationStrategy interface {
	// ExtractValue extracts the value to aggregate from an assignment
	ExtractValue(trip *dto.EmployeeTripDTO) int
	// ExtractValueFromBusinessTrip extracts value from business trip (for trip count)
	ExtractValueFromBusinessTrip(trip *dto.BuisnessTripDTO) int
	// GetDataSource returns what data to fetch - "all_employees" or "all_trips"
	GetDataSource() string
}

// MoneySpentStrategy aggregates total money spent per year
type MoneySpentStrategy struct{}

func (m *MoneySpentStrategy) ExtractValue(trip *dto.EmployeeTripDTO) int {
	return trip.MoneySpent
}

func (m *MoneySpentStrategy) ExtractValueFromBusinessTrip(trip *dto.BuisnessTripDTO) int {
	return 0 // Not used for business trip aggregation
}

func (m *MoneySpentStrategy) GetDataSource() string {
	return "all_employees"
}

// TripCountStrategy counts trips per year
type TripCountStrategy struct{}

func (t *TripCountStrategy) ExtractValue(trip *dto.EmployeeTripDTO) int {
	return 1 // Count each trip as 1
}

func (t *TripCountStrategy) ExtractValueFromBusinessTrip(trip *dto.BuisnessTripDTO) int {
	return 1 // Count each trip as 1
}

func (t *TripCountStrategy) GetDataSource() string {
	return "all_trips"
}

// EmployeeStatStrategy aggregates both trip count and money spent for statistics
type EmployeeStatStrategy struct{}

func (e *EmployeeStatStrategy) ExtractValue(trip *dto.EmployeeTripDTO) int {
	return trip.MoneySpent
}

func (e *EmployeeStatStrategy) ExtractValueFromBusinessTrip(trip *dto.BuisnessTripDTO) int {
	return 0
}

func (e *EmployeeStatStrategy) GetDataSource() string {
	return "single_employee"
}
