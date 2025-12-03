# Service Layer Architecture

## Component Diagram

```
┌─────────────────────────────────────────────────────────────┐
│                    HTTP Controllers                          │
│         (main_controller, employee_controller)              │
└────────────────────────┬──────────────────────────────────┘
                         │
                    calls methods
                         │
        ┌────────────────▼─────────────────┐
        │     Service Layer               │
        │  (business logic orchestration)  │
        └────────────────┬─────────────────┘
                         │
         ┌───────────────┴───────────────┐
         │                               │
    ┌────▼───────────────┐     ┌────────▼──────────────┐
    │  Template Methods  │     │   Strategy Pattern   │
    │  (algorithms)      │     │   (value extraction) │
    │                    │     │                      │
    │ • aggregateBy      │────▶│ • MoneySpentStrategy │
    │   YearsWithStrategy│     │ • TripCountStrategy │
    │                    │     │ • EmployeeStatStrat  │
    │ • aggregateTrips   │     └───────────────────────┘
    │   WithStrategy     │
    └────────┬───────────┘
             │
             │ uses
             │
    ┌────────▼──────────────────┐
    │  Aggregator Classes       │
    │  (data structures)        │
    │                           │
    │ • YearlyAggregator        │
    │   - Manages: map[year]val │
    │   - Sorts results         │
    │                           │
    │ • YearlyStatAggregator    │
    │   - Manages dual counters │
    │   - Calculates averages   │
    └────────┬──────────────────┘
             │
             │ fetches data via
             │
    ┌────────▼──────────────────┐
    │  Repository Layer         │
    │                           │
    │ • EmployeeRepo            │
    │ • BusinessTripRepo        │
    └───────────────────────────┘
```

## Class Relationships

### Strategy Pattern Hierarchy

```
AggregationStrategy (interface)
├── MoneySpentStrategy
│   └── ExtractValue(): returns trip.MoneySpent
│   └── ExtractValueFromBusinessTrip(): returns 0
│
├── TripCountStrategy
│   └── ExtractValue(): returns 1
│   └── ExtractValueFromBusinessTrip(): returns 1
│
└── EmployeeStatStrategy
    └── ExtractValue(): returns trip.MoneySpent
    └── ExtractValueFromBusinessTrip(): returns 0
```

### Template Method Pattern

```
Service (uses template methods)
├── aggregateByYearsWithStrategy(strategy)
│   └── Orchestrates the algorithm:
│       1. Fetch from employeeRepo
│       2. Iterate employees → trips
│       3. Strategy.ExtractValue(trip)
│       4. Aggregator.AddValue(year, value)
│       5. Return aggregator.GetResults()
│
└── aggregateTripsWithStrategy(strategy)
    └── Orchestrates the algorithm:
        1. Fetch from businessTripRepo
        2. Iterate trips
        3. Strategy.ExtractValueFromBusinessTrip(trip)
        4. Aggregator.AddValue(year, value)
        5. Return aggregator.GetResults()
```

### Aggregator Hierarchy

```
Aggregators (Data Structures)
├── YearlyAggregator
│   ├── State: map[int]int (year → value)
│   ├── AddValue(year, value)
│   └── GetResults() → []GraphData (sorted)
│
└── YearlyStatAggregator
    ├── State: map[int]int (trips), map[int]int (money)
    ├── AddValue(year, tripCount, moneySpent)
    ├── GetTotalTripCount() → int
    ├── GetTotalMoneySpent() → int
    ├── GetAverageTripsPerYear() → float32
    └── GetAverageMoneyPerYear() → float32
```

## Data Flow Examples

### Example 1: GetMoneySpentByAllYears()

```
User calls: service.GetMoneySpentByAllYears()
              │
              ▼
        Create MoneySpentStrategy
              │
              ▼
        Call aggregateByYearsWithStrategy(strategy)
              │
              ├─► employeeRepo.All()
              │
              ├─► For each employee:
              │   └─► For each trip:
              │       ├─► strategy.ExtractValue(trip)
              │       │   └─► Returns trip.MoneySpent
              │       │
              │       └─► aggregator.AddValue(year, money)
              │
              └─► aggregator.GetResults()
                  ├─► Builds GraphData from map
                  └─► Sorts by year
                      └─► Returns []GraphData
```

### Example 2: GetEmployeeStat(id)

```
User calls: service.GetEmployeeStat(id)
              │
              ▼
        employeeRepo.Find(id)
              │
              ▼
        Create YearlyStatAggregator
              │
              ▼
        For each trip:
          └─► aggregator.AddValue(year, 1, money)
              ├─► tripCounters[year] += 1
              └─► moneyCounters[year] += money
              │
              ▼
        Build EmployeeData:
          ├─► name: from employee
          ├─► tripCount: aggregator.GetTotalTripCount()
          ├─► moneySpent: aggregator.GetTotalMoneySpent()
          ├─► avgTripCount: aggregator.GetAverageTripsPerYear()
          └─► avgMoneySpent: aggregator.GetAverageMoneyPerYear()
```

## Dependency Injection

```
main() 
  │
  ├─► Create repos:
  │   ├── employeeRepo := employee_repo.New(db)
  │   └── businessTripRepo := business_trip_repo.New(db)
  │
  ├─► Create service:
  │   └── service := service.New(employeeRepo, businessTripRepo)
  │
  ├─► Create controllers:
  │   ├── mainCtrl := main_controller.New(service)
  │   └── empCtrl := employee_controller.New(service)
  │
  └─► Register routes with controllers
```

## Separation of Concerns

| Layer | Responsibility | Changes When |
|-------|---|---|
| **Strategies** | Extract values from data | New aggregation types needed |
| **Aggregators** | Store and transform data | Different time periods (yearly → monthly) |
| **Templates** | Orchestrate algorithm | Overall aggregation flow changes |
| **Service** | Public API, delegation | New public operations needed |
| **Controllers** | HTTP handling | API endpoints change |
| **Repos** | Data access | Schema or database changes |

## Adding New Functionality

### Scenario: Add "Get Trips by Department"

**Step 1**: Create strategy
```go
type DepartmentTripCountStrategy struct{}

func (d *DepartmentTripCountStrategy) ExtractValue(trip *dto.EmployeeTripDTO) int {
    return 1  // Count trips
}
```

**Step 2**: Modify aggregator (or create new one)
```go
type DepartmentAggregator struct {
    counters map[string]int  // department → count
}
```

**Step 3**: Add template method
```go
func (s *Service) aggregateByDepartmentWithStrategy(strategy DepartmentStrategy) map[string]int {
    // Similar to aggregateByYears but groups by department
}
```

**Step 4**: Add public method
```go
func (s *Service) GetTripsByDepartment() map[string]int {
    strategy := &DepartmentTripCountStrategy{}
    return s.aggregateByDepartmentWithStrategy(strategy)
}
```

**Step 5**: Use in controller
```go
func (c *Controller) GetDepartmentStats(w http.ResponseWriter, r *http.Request) {
    data := c.service.GetTripsByDepartment()
    json.NewEncoder(w).Encode(data)
}
```

## Pattern Selection Guide

| Problem | Pattern | Solution |
|---------|---------|----------|
| Multiple ways to extract values | Strategy | Create strategy implementations |
| Duplicate iteration/sorting logic | Template Method | Create aggregator template |
| New aggregation types | Both | New strategy + use existing template |
| Different time groupings | Template Method | New aggregator implementation |
| Complex filtering | Strategy | Filtering strategy decorator |

## Performance Characteristics

| Operation | Time | Space | Notes |
|-----------|------|-------|-------|
| GetMoneySpentByAllYears() | O(n) | O(y) | n=trips, y=years |
| GetEmployeeStat() | O(t) | O(y) | t=employee's trips |
| aggregateByYears...() | O(n) | O(y) | Template overhead negligible |
| Strategy.ExtractValue() | O(1) | O(1) | Single field access |
| Aggregator.AddValue() | O(1) | - | Map lookup |
| Aggregator.GetResults() | O(y log y) | O(y) | Sorting cost |

---

See `DESIGN_PATTERNS.md` for detailed pattern explanations and `EXAMPLES.md` for usage examples.
