# Service Layer Examples

## How to Use the Strategy Pattern

### Example 1: Add a New Aggregation Strategy

If you want to add a new aggregation type (e.g., count destinations by year):

```go
// 1. Create a new strategy in strategies.go
type DestinationCountStrategy struct{}

func (d *DestinationCountStrategy) ExtractValue(trip *dto.EmployeeTripDTO) int {
    return 1 // Count each unique destination
}

func (d *DestinationCountStrategy) ExtractValueFromBusinessTrip(trip *dto.BuisnessTripDTO) int {
    return 1
}

func (d *DestinationCountStrategy) GetDataSource() string {
    return "all_trips"
}

// 2. Add a public method to Service
func (s *Service) GetDestinationCountByYears() *[]GraphData {
    strategy := &DestinationCountStrategy{}
    return s.aggregateTripsWithStrategy(strategy)
}

// 3. Use it in controller
func (c *Controller) GetDestinationData(w http.ResponseWriter, r *http.Request) {
    data := c.service.GetDestinationCountByYears()
    // ... render to JSON/HTML
}
```

### Example 2: Modify Existing Strategy

To change how money spent is calculated (e.g., convert to different currency):

```go
// Update MoneySpentStrategy
type MoneySpentStrategy struct {
    conversionRate float64
}

func (m *MoneySpentStrategy) ExtractValue(trip *dto.EmployeeTripDTO) int {
    return int(float64(trip.MoneySpent) * m.conversionRate)
}

// Use it
func (s *Service) GetMoneySpentByAllYearsInEuros() *[]GraphData {
    strategy := &MoneySpentStrategy{conversionRate: 0.92}
    return s.aggregateByYearsWithStrategy(strategy)
}
```

---

## How to Use the Template Method Pattern

### Example 3: Custom Aggregator

If you need different aggregation behavior (not yearly), create a custom aggregator:

```go
// Create a new aggregator for quarterly data
type QuarterlyAggregator struct {
    counters map[string]int // "2023-Q1", "2023-Q2", etc.
}

func NewQuarterlyAggregator() *QuarterlyAggregator {
    return &QuarterlyAggregator{
        counters: make(map[string]int),
    }
}

func (qa *QuarterlyAggregator) AddValue(year int, month int, value int) {
    quarter := fmt.Sprintf("%d-Q%d", year, (month-1)/3+1)
    qa.counters[quarter] += value
}

func (qa *QuarterlyAggregator) GetResults() interface{} {
    // Return quarterly data
    return qa.counters
}

// Use with Service
func (s *Service) GetMoneySpentByQuarters() interface{} {
    data, _ := s.employeeRepo.All()
    aggregator := NewQuarterlyAggregator()
    
    for _, employee := range *data {
        for _, trip := range employee.Trips {
            year := trip.BuisnessTrip.StartAt.Year()
            month := int(trip.BuisnessTrip.StartAt.Month())
            aggregator.AddValue(year, month, trip.MoneySpent)
        }
    }
    
    return aggregator.GetResults()
}
```

---

## Real-World Usage Examples

### Example 4: Service Method Simplification

**Before**:
```go
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
        res = append(res, GraphData{X: k, Y: v})
    }
    
    sort.Slice(res, func(i, j int) bool {
        return res[i].X < res[j].X
    })
    
    return &res
}
```

**After**:
```go
func (s *Service) GetMoneySpentByAllYears() *[]GraphData {
    strategy := &MoneySpentStrategy{}
    return s.aggregateByYearsWithStrategy(strategy)
}
```

### Example 5: Testing with Mock Strategy

```go
type MockStrategy struct {
    values []int
}

func (m *MockStrategy) ExtractValue(trip *dto.EmployeeTripDTO) int {
    if len(m.values) > 0 {
        val := m.values[0]
        m.values = m.values[1:]
        return val
    }
    return 0
}

// In your test
func TestAggregation(t *testing.T) {
    service := &Service{...}
    mockStrategy := &MockStrategy{values: []int{100, 200}}
    
    // Call template method directly
    results := service.aggregateByYearsWithStrategy(mockStrategy)
    
    // Verify results
    assert.Len(t, *results, 1)
}
```

---

## When to Use Each Pattern

### Use Strategy When:
- You have multiple similar operations that differ in their core logic
- You want to avoid huge if/else chains in methods
- You expect new variations to be added frequently
- You want to enable runtime selection of behavior

### Use Template Method When:
- Multiple methods share the same algorithmic structure
- You have complex algorithms that are repeated across methods
- You want to enforce a specific sequence of operations
- You want subclasses (or in our case, different data aggregators) to hook into specific steps

### Combine Them When:
- The overall algorithm is fixed but internal steps vary
- You have multiple variations of both the algorithm AND the data extraction
- You want maximum flexibility with minimal code duplication
