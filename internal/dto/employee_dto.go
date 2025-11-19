package dto

type EmployeeDTO struct {
	ID   uint
	Name string
	Trips []EmployeeTripDTO
}
