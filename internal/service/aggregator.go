package service

import (
	"sort"
)

type YearlyAggregator struct {
	counters map[int]int
}

func NewYearlyAggregator() *YearlyAggregator {
	return &YearlyAggregator{
		counters: make(map[int]int),
	}
}

func (ya *YearlyAggregator) AddValue(year int, value int) {
	if _, exists := ya.counters[year]; !exists {
		ya.counters[year] = value
	} else {
		ya.counters[year] += value
	}
}

func (ya *YearlyAggregator) GetResults() *[]GraphData {
	res := []GraphData{}

	for k, v := range ya.counters {
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

type YearlyStatAggregator struct {
	tripCounters  map[int]int
	moneyCounters map[int]int
}

func NewYearlyStatAggregator() *YearlyStatAggregator {
	return &YearlyStatAggregator{
		tripCounters:  make(map[int]int),
		moneyCounters: make(map[int]int),
	}
}

func (ysa *YearlyStatAggregator) AddValue(year int, tripCount int, moneySpent int) {
	if _, exists := ysa.tripCounters[year]; !exists {
		ysa.tripCounters[year] = tripCount
	} else {
		ysa.tripCounters[year] += tripCount
	}

	if _, exists := ysa.moneyCounters[year]; !exists {
		ysa.moneyCounters[year] = moneySpent
	} else {
		ysa.moneyCounters[year] += moneySpent
	}
}

func (ysa *YearlyStatAggregator) GetTotalTripCount() int {
	sum := 0
	for _, v := range ysa.tripCounters {
		sum += v
	}
	return sum
}

func (ysa *YearlyStatAggregator) GetTotalMoneySpent() int {
	sum := 0
	for _, v := range ysa.moneyCounters {
		sum += v
	}
	return sum
}

func (ysa *YearlyStatAggregator) GetAverageTripsPerYear() float32 {
	if len(ysa.tripCounters) == 0 {
		return 0
	}
	return float32(ysa.GetTotalTripCount()) / float32(len(ysa.tripCounters))
}

func (ysa *YearlyStatAggregator) GetAverageMoneyPerYear() float32 {
	if len(ysa.moneyCounters) == 0 {
		return 0
	}
	return float32(ysa.GetTotalMoneySpent()) / float32(len(ysa.moneyCounters))
}
