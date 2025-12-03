package service_test

import (
	"TP_Andreev/internal/dto"
	"TP_Andreev/internal/service"
	"slices"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
)

type mockEmployeeRepo struct {
	mock.Mock
}

func (m *mockEmployeeRepo) Find(id uint) (*dto.EmployeeDTO, error) {
	args := m.Called(id)
	return args.Get(0).(*dto.EmployeeDTO), args.Error(1)
}

func (m *mockEmployeeRepo) All() (*[]dto.EmployeeDTO, error) {
	args := m.Called()
	return args.Get(0).(*[]dto.EmployeeDTO), args.Error(1)
}

type mockBusinessTripRepo struct {
	mock.Mock
}

func (m *mockBusinessTripRepo) All() (*[]dto.BuisnessTripDTO, error) {
	args := m.Called()
	return args.Get(0).(*[]dto.BuisnessTripDTO), args.Error(1)
}

var employeeDtoArray *[]dto.EmployeeDTO = &[]dto.EmployeeDTO{
	{
		ID:   1,
		Name: "A",
		Trips: []dto.EmployeeTripDTO{
			{
				MoneySpent: 10,
				BuisnessTrip: dto.BuisnessTripDTO{
					Destination: "Dest1",
					StartAt:     time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
					EndAt:       time.Date(2020, 1, 12, 0, 0, 0, 0, time.UTC),
				},
			},
			{
				MoneySpent: 20,
				BuisnessTrip: dto.BuisnessTripDTO{
					Destination: "Dest2",
					StartAt:     time.Date(2021, 2, 21, 0, 0, 0, 0, time.UTC),
					EndAt:       time.Date(2021, 2, 25, 0, 0, 0, 0, time.UTC),
				},
			},
		},
	},
	{
		ID:   2,
		Name: "B",
		Trips: []dto.EmployeeTripDTO{
			{
				MoneySpent: 5,
				BuisnessTrip: dto.BuisnessTripDTO{
					Destination: "Dest2",
					StartAt:     time.Date(2021, 2, 21, 0, 0, 0, 0, time.UTC),
					EndAt:       time.Date(2021, 2, 25, 0, 0, 0, 0, time.UTC),
				},
			},
			{
				MoneySpent: 15,
				BuisnessTrip: dto.BuisnessTripDTO{
					Destination: "Dest3",
					StartAt:     time.Date(2022, 3, 5, 0, 0, 0, 0, time.UTC),
					EndAt:       time.Date(2022, 3, 10, 0, 0, 0, 0, time.UTC),
				},
			},
		},
	},
}

var employeeDto *dto.EmployeeDTO = &dto.EmployeeDTO{
	Name: "Name",
	Trips: []dto.EmployeeTripDTO{
		{
			MoneySpent:   10,
			BuisnessTrip: dto.BuisnessTripDTO{StartAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)},
		},
		{
			MoneySpent:   10,
			BuisnessTrip: dto.BuisnessTripDTO{StartAt: time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)},
		},
		{
			MoneySpent:   10,
			BuisnessTrip: dto.BuisnessTripDTO{StartAt: time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC)},
		},
		{
			MoneySpent:   10,
			BuisnessTrip: dto.BuisnessTripDTO{StartAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
		},
		{
			MoneySpent:   10,
			BuisnessTrip: dto.BuisnessTripDTO{StartAt: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)},
		},
		{
			MoneySpent:   10,
			BuisnessTrip: dto.BuisnessTripDTO{StartAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)},
		},
	},
}

var buisnessTripDtoArray *[]dto.BuisnessTripDTO = &[]dto.BuisnessTripDTO{
	{ID: 1, StartAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)},
	{ID: 2, StartAt: time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)},
	{ID: 3, StartAt: time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC)},
	{ID: 4, StartAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
	{ID: 5, StartAt: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)},
	{ID: 6, StartAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)},
}

func TestGetAllEmployeeTrips(t *testing.T) {
	mockEmployeeRepo := new(mockEmployeeRepo)
	mockBusinessTripRepo := new(mockBusinessTripRepo)

	mockEmployeeRepo.On("All").Return(
		employeeDtoArray,
		nil,
	)

	expected := []service.EmployeeTripData{
		{Id: 2, Name: "B", Destination: "Dest3", Date: "05.03.2022", Duration: 5, MoneySpent: 15},
		{Id: 1, Name: "A", Destination: "Dest2", Date: "21.02.2021", Duration: 4, MoneySpent: 20},
		{Id: 2, Name: "B", Destination: "Dest2", Date: "21.02.2021", Duration: 4, MoneySpent: 5},
		{Id: 1, Name: "A", Destination: "Dest1", Date: "01.01.2020", Duration: 11, MoneySpent: 10},
	}

	service := service.New(mockEmployeeRepo, mockBusinessTripRepo)

	actual := service.GetAllEmployeeTrips()

	if !slices.Equal(expected, *actual) {
		t.Errorf("Result was incorrect, got: %v, want: %v.", *actual, expected)
	}
}

func TestGetMoneySpentByAllYears(t *testing.T) {
	mockEmployeeRepo := new(mockEmployeeRepo)
	mockBusinessTripRepo := new(mockBusinessTripRepo)

	mockEmployeeRepo.On("All").Return(
		employeeDtoArray,
		nil,
	)

	expected := []service.GraphData{
		{X: 2020, Y: 10},
		{X: 2021, Y: 25},
		{X: 2022, Y: 15},
	}

	service := service.New(mockEmployeeRepo, mockBusinessTripRepo)

	actual := service.GetMoneySpentByAllYears()

	if !slices.Equal(expected, *actual) {
		t.Errorf("Result was incorrect, got: %v, want: %v.", *actual, expected)
	}
}

func TestGetTripCountByAllYears(t *testing.T) {
	mockEmployeeRepo := new(mockEmployeeRepo)
	mockBusinessTripRepo := new(mockBusinessTripRepo)

	mockBusinessTripRepo.On("All").Return(
		buisnessTripDtoArray,
		nil,
	)

	expected := []service.GraphData{
		{X: 2020, Y: 3},
		{X: 2021, Y: 2},
		{X: 2022, Y: 1},
	}

	service := service.New(mockEmployeeRepo, mockBusinessTripRepo)

	actual := service.GetTripCountByAllYears()

	if !slices.Equal(expected, *actual) {
		t.Errorf("Result was incorrect, got: %v, want: %v.", *actual, expected)
	}
}

func TestGetEmployeeTripCountByAllYears(t *testing.T) {
	mockEmployeeRepo := new(mockEmployeeRepo)
	mockBusinessTripRepo := new(mockBusinessTripRepo)

	mockEmployeeRepo.On("Find", uint(1)).Return(
		employeeDto,
		nil,
	)

	expected := []service.GraphData{
		{X: 2020, Y: 3},
		{X: 2021, Y: 2},
		{X: 2022, Y: 1},
	}

	service := service.New(mockEmployeeRepo, mockBusinessTripRepo)

	actual := service.GetEmployeeTripCountByAllYears(1)

	if !slices.Equal(expected, *actual) {
		t.Errorf("Result was incorrect, got: %v, want: %v.", *actual, expected)
	}
}

func TestGetEmployeeStat(t *testing.T) {
	mockEmployeeRepo := new(mockEmployeeRepo)
	mockBusinessTripRepo := new(mockBusinessTripRepo)

	mockEmployeeRepo.On("Find", uint(1)).Return(
		employeeDto,
		nil,
	)

	expected := service.EmployeeData{
		Name:          "Name",
		TripCount:     6,
		MoneySpent:    60,
		AvgTripCount:  2,
		AvgMoneySpent: 20,
	}

	service := service.New(mockEmployeeRepo, mockBusinessTripRepo)

	actual := service.GetEmployeeStat(1)

	if *actual != expected {
		t.Errorf("Result was incorrect, got: %v, want: %v.", *actual, expected)
	}
}
