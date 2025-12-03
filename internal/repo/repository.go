package repository

import "TP_Andreev/internal/dto"

type EmployeeRepo interface {
	Find(id uint) (*dto.EmployeeDTO, error)
	All() (*[]dto.EmployeeDTO, error)
}

type BusinessTripRepo interface {
	All() (*[]dto.BuisnessTripDTO, error)
}
