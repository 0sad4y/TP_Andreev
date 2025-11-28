package service

import (
	"TP_Andreev/internal/repo"
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
			date := t.BuisnessTrip.StartAt.Format("01.02.2006")
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

	// return &[]EmployeeTripData{
	// 	{Id: 1, Name: "Имя 1", Date: "01.01.2025", Duration: 2, Destination: "Куда A", MoneySpent: 500},
	// 	{Id: 2, Name: "Имя 2", Date: "05.01.2025", Duration: 3, Destination: "Куда B", MoneySpent: 800},
	// 	{Id: 3, Name: "Имя 3", Date: "10.01.2025", Duration: 1, Destination: "Куда C", MoneySpent: 300},
	// 	{Id: 4, Name: "Имя 4", Date: "12.01.2025", Duration: 4, Destination: "Куда D", MoneySpent: 700},
	// 	{Id: 5, Name: "Имя 5", Date: "15.01.2025", Duration: 5, Destination: "Куда E", MoneySpent: 1200},
	// 	{Id: 6, Name: "Имя 6", Date: "18.01.2025", Duration: 2, Destination: "Куда F", MoneySpent: 400},
	// 	{Id: 7, Name: "Имя 7", Date: "20.01.2025", Duration: 3, Destination: "Куда G", MoneySpent: 900},
	// 	{Id: 8, Name: "Имя 8", Date: "22.01.2025", Duration: 2, Destination: "Куда H", MoneySpent: 600},
	// 	{Id: 9, Name: "Имя 9", Date: "25.01.2025", Duration: 1, Destination: "Куда I", MoneySpent: 250},
	// 	{Id: 10, Name: "Имя 10", Date: "28.01.2025", Duration: 4, Destination: "Куда J", MoneySpent: 1000},
	// }
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

	return &res

	// return &[]GraphData{
	// 	{X: 2001, Y: 10000},
	// 	{X: 2002, Y: 11000},
	// 	{X: 2003, Y: 12000},
	// 	{X: 2004, Y: 13000},
	// 	{X: 2005, Y: 20000},
	// 	{X: 2006, Y: 25000},
	// 	{X: 2007, Y: 30000},
	// 	{X: 2008, Y: 10000},
	// 	{X: 2009, Y: 12000},
	// 	{X: 2010, Y: 23000},
	// }
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

	return &res

	// return &[]GraphData{
	// 	{X: 2001, Y: 3},
	// 	{X: 2002, Y: 4},
	// 	{X: 2003, Y: 3},
	// 	{X: 2004, Y: 4},
	// 	{X: 2005, Y: 5},
	// 	{X: 2006, Y: 6},
	// 	{X: 2007, Y: 7},
	// 	{X: 2008, Y: 3},
	// 	{X: 2009, Y: 4},
	// 	{X: 2010, Y: 5},
	// }
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

	return &res

	// return &[]GraphData{
	// 	{X: 2001, Y: 10000},
	// 	{X: 2002, Y: 11000},
	// 	{X: 2003, Y: 12000},
	// 	{X: 2004, Y: 13000},
	// 	{X: 2005, Y: 20000},
	// 	{X: 2006, Y: 25000},
	// 	{X: 2007, Y: 30000},
	// 	{X: 2008, Y: 10000},
	// 	{X: 2009, Y: 12000},
	// 	{X: 2010, Y: 23000},
	// }
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
