package service

import (
	"TP_Andreev/internal/dto"
)

type AggregationStrategy interface {
	ExtractValue(trip *dto.EmployeeTripDTO) int
	ExtractValueFromBusinessTrip(trip *dto.BuisnessTripDTO) int
}

type MoneySpentStrategy struct{}

func (m *MoneySpentStrategy) ExtractValue(trip *dto.EmployeeTripDTO) int {
	return trip.MoneySpent
}

func (m *MoneySpentStrategy) ExtractValueFromBusinessTrip(trip *dto.BuisnessTripDTO) int {
	return 0
}

type TripCountStrategy struct{}

func (t *TripCountStrategy) ExtractValue(trip *dto.EmployeeTripDTO) int {
	return 1
}

func (t *TripCountStrategy) ExtractValueFromBusinessTrip(trip *dto.BuisnessTripDTO) int {
	return 1
}
