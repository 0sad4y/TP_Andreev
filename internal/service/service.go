package service

import (
	"TP_Andreev/internal/repo"
	"sort"
)

type Service struct {
	employeeRepo     repo.EmployeeRepo
	businessTripRepo repo.BusinessTripRepo
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

func New(employeeRepo repo.EmployeeRepo, businessTripRepo repo.BusinessTripRepo) *Service {
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

	return &res
}

func (s *Service) GetMoneySpentByAllYears() *[]GraphData {
	data, _ := s.employeeRepo.All()

	counter := map[int]int{}

	for _, d := range *data {
		for _, t := range d.Trips {
			year := t.BuisnessTrip.StartAt.Year()
			moneySpent := t.MoneySpent

			_, ok := counter[year]
			if !ok {
				counter[year] = moneySpent
			} else {
				counter[year] += moneySpent
			}
		}
	}

	res := []GraphData{}

	for k, v := range counter {
		data := GraphData{
			X: k,
			Y: v,
		}
		res = append(res, data)
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].X < res[j].X
	})

	return &res
}

func (s *Service) GetTripCountByAllYears() *[]GraphData {
	data, _ := s.businessTripRepo.All()

	counter := map[int]int{}

	for _, d := range *data {
		year := d.StartAt.Year()

		_, ok := counter[year]
		if !ok {
			counter[year] = 1
		} else {
			counter[year]++
		}
	}

	res := []GraphData{}

	for k, v := range counter {
		data := GraphData{
			X: k,
			Y: v,
		}
		res = append(res, data)
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].X < res[j].X
	})

	return &res
}

func (s *Service) GetEmployeeTripCountByAllYears(id int) *[]GraphData {
	data, _ := s.employeeRepo.Find(uint(id))

	counter := map[int]int{}

	for _, t := range (*data).Trips {
		year := t.BuisnessTrip.StartAt.Year()

		_, ok := counter[year]
		if !ok {
			counter[year] = 1
		} else {
			counter[year]++
		}
	}

	res := []GraphData{}

	for k, v := range counter {
		data := GraphData{
			X: k,
			Y: v,
		}
		res = append(res, data)
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].X < res[j].X
	})

	return &res
}

func (s *Service) GetEmployeeStat(id int) *EmployeeData {
	data, _ := s.employeeRepo.Find(uint(id))
	name := data.Name

	tripCounter := map[int]int{}
	moneyCounter := map[int]int{}

	for _, t := range (*data).Trips {
		year := t.BuisnessTrip.StartAt.Year()
		moneySpent := t.MoneySpent

		_, ok := tripCounter[year]
		if !ok {
			tripCounter[year] = 1
		} else {
			tripCounter[year]++
		}

		_, ok = moneyCounter[year]
		if !ok {
			moneyCounter[year] = moneySpent
		} else {
			moneyCounter[year] += moneySpent
		}
	}

	tripSum := 0
	moneySum := 0

	for _, v := range tripCounter {
		tripSum += v
	}
	for _, v := range moneyCounter {
		moneySum += v
	}

	return &EmployeeData{
		Name:          name,
		TripCount:     tripSum,
		MoneySpent:    moneySum,
		AvgTripCount:  float32(tripSum) / float32(len(tripCounter)),
		AvgMoneySpent: float32(moneySum) / float32(len(moneyCounter)),
	}
}
