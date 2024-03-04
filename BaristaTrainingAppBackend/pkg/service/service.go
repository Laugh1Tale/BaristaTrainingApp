package service

import (
	"barTrApp/pkg/auth"
	"barTrApp/pkg/repository"
)

type Authorization interface {
	CreateEmployee(employee auth.Employee) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Authorization: newAuthService(repository.Authorization),
	}
}
