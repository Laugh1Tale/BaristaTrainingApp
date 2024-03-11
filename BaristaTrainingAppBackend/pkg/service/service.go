package service

import (
	"barTrApp/pkg/model"
	"barTrApp/pkg/repository"
)

type Authorization interface {
	CreateEmployee(employee model.Employee) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Authorization: newAuthService(repository.Authorization),
	}
}
