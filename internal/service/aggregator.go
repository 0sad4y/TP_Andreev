package service

import (
	"sort"
)

// YearlyAggregator implements Template Method pattern for yearly aggregation
type YearlyAggregator struct {
	counters map[int]int
}

// NewYearlyAggregator creates a new aggregator
func NewYearlyAggregator() *YearlyAggregator {
	return &YearlyAggregator{
		counters: make(map[int]int),
	}
}

// AddValue adds a value to the counter for a given year (Template Method)
func (ya *YearlyAggregator) AddValue(year int, value int) {
	if _, exists := ya.counters[year]; !exists {
		ya.counters[year] = value
	} else {
		ya.counters[year] += value
	}
}

// GetResults returns sorted GraphData results
func (ya *YearlyAggregator) GetResults() *[]GraphData {
	res := []GraphData{}

	for k, v := range ya.counters {
		data := GraphData{
			X: k,
			Y: v,
		}
		res = append(res, data)
	}

	// Sort by year
	sort.Slice(res, func(i, j int) bool {
		return res[i].X < res[j].X
	})

	return &res
}

// YearlyStatAggregator aggregates both trip count and money spent statistics per year
type YearlyStatAggregator struct {
	tripCounters map[int]int
	moneyCounters map[int]int
}

// NewYearlyStatAggregator creates a new stat aggregator
func NewYearlyStatAggregator() *YearlyStatAggregator {
	return &YearlyStatAggregator{
		tripCounters: make(map[int]int),
		moneyCounters: make(map[int]int),
	}
}

// AddValue adds both trip count and money spent for a year (Template Method pattern)
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

// GetTotalTripCount returns total trips across all years
func (ysa *YearlyStatAggregator) GetTotalTripCount() int {
	sum := 0
	for _, v := range ysa.tripCounters {
		sum += v
	}
	return sum
}

// GetTotalMoneySpent returns total money spent across all years
func (ysa *YearlyStatAggregator) GetTotalMoneySpent() int {
	sum := 0
	for _, v := range ysa.moneyCounters {
		sum += v
	}
	return sum
}

// GetAverageTripsPerYear returns average trips per year
func (ysa *YearlyStatAggregator) GetAverageTripsPerYear() float32 {
	if len(ysa.tripCounters) == 0 {
		return 0
	}
	return float32(ysa.GetTotalTripCount()) / float32(len(ysa.tripCounters))
}

// GetAverageMoneyPerYear returns average money spent per year
func (ysa *YearlyStatAggregator) GetAverageMoneyPerYear() float32 {
	if len(ysa.moneyCounters) == 0 {
		return 0
	}
	return float32(ysa.GetTotalMoneySpent()) / float32(len(ysa.moneyCounters))
}
