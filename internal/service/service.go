package service

import (
	repository "TP_Andreev/internal/repo"
	"sort"
	"time"
)

type Service struct {
	employeeRepo     repository.EmployeeRepo
	businessTripRepo repository.BusinessTripRepo
}

type EmployeeTripData struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Destination string `json:"destination"`
	Date        string `json:"date"`
	Duration    int    `json:"duration"`
	MoneySpent  int    `json:"moneySpent"`
}

type EmployeeData struct {
	Name          string  `json:"name"`
	TripCount     int     `json:"tripCount"`
	MoneySpent    int     `json:"moneySpent"`
	AvgTripCount  float32 `json:"avgTripCount"`
	AvgMoneySpent float32 `json:"avgMoneySpent"`
}

type GraphData struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func New(employeeRepo repository.EmployeeRepo, businessTripRepo repository.BusinessTripRepo) *Service {
	return &Service{employeeRepo: employeeRepo, businessTripRepo: businessTripRepo}
}

func (s *Service) GetAllEmployeeTrips() *[]EmployeeTripData {
	data, _ := s.employeeRepo.All()

	res := []EmployeeTripData{}
	for _, d := range *data {
		id := d.ID
		name := d.Name
		for _, t := range d.Trips {
			date := t.BuisnessTrip.StartAt.Format("02.01.2006")
			duration := int(t.BuisnessTrip.EndAt.Sub(t.BuisnessTrip.StartAt).Hours()) / 24
			destination := t.BuisnessTrip.Destination
			moneySpent := t.MoneySpent

			empTripData := EmployeeTripData{
				Id:          id,
				Name:        name,
				Date:        date,
				Duration:    duration,
				Destination: destination,
				MoneySpent:  moneySpent,
			}

			res = append(res, empTripData)
		}
	}

	sort.Slice(res, func(i, j int) bool {
		t1, _ := time.Parse("02.01.2006", res[i].Date)
		t2, _ := time.Parse("02.01.2006", res[j].Date)
		return t2.Before(t1)
	})

	return &res
}

func (s *Service) GetMoneySpentByAllYears() *[]GraphData {
	strategy := &MoneySpentStrategy{}
	return s.aggregateByYearsWithStrategy(strategy)
}

func (s *Service) GetTripCountByAllYears() *[]GraphData {
	strategy := &TripCountStrategy{}
	return s.aggregateTripsWithStrategy(strategy)
}

func (s *Service) GetEmployeeTripCountByAllYears(id int) *[]GraphData {
	data, _ := s.employeeRepo.Find(uint(id))

	aggregator := NewYearlyAggregator()
	for _, t := range (*data).Trips {
		year := t.BuisnessTrip.StartAt.Year()
		aggregator.AddValue(year, 1)
	}

	return aggregator.GetResults()
}

func (s *Service) GetEmployeeStat(id int) *EmployeeData {
	data, _ := s.employeeRepo.Find(uint(id))
	name := data.Name

	aggregator := NewYearlyStatAggregator()
	for _, t := range (*data).Trips {
		year := t.BuisnessTrip.StartAt.Year()
		aggregator.AddValue(year, 1, t.MoneySpent)
	}

	return &EmployeeData{
		Name:          name,
		TripCount:     aggregator.GetTotalTripCount(),
		MoneySpent:    aggregator.GetTotalMoneySpent(),
		AvgTripCount:  aggregator.GetAverageTripsPerYear(),
		AvgMoneySpent: aggregator.GetAverageMoneyPerYear(),
	}
}

func (s *Service) aggregateByYearsWithStrategy(strategy AggregationStrategy) *[]GraphData {
	data, _ := s.employeeRepo.All()
	aggregator := NewYearlyAggregator()

	for _, employee := range *data {
		for _, trip := range employee.Trips {
			year := trip.BuisnessTrip.StartAt.Year()
			value := strategy.ExtractValue(&trip)
			aggregator.AddValue(year, value)
		}
	}

	return aggregator.GetResults()
}

func (s *Service) aggregateTripsWithStrategy(strategy AggregationStrategy) *[]GraphData {
	data, _ := s.businessTripRepo.All()
	aggregator := NewYearlyAggregator()

	for _, trip := range *data {
		year := trip.StartAt.Year()
		value := strategy.ExtractValueFromBusinessTrip(&trip)
		aggregator.AddValue(year, value)
	}

	return aggregator.GetResults()
}
