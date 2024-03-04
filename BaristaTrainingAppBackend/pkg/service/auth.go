package service

import (
	"barTrApp/pkg/auth"
	"barTrApp/pkg/repository"
	"crypto/sha1"
	"fmt"
)

const salt = "kahj7832jkasdfh23"

type AuthService struct {
	repository repository.Authorization
}

func newAuthService(repository repository.Authorization) *AuthService {
	return &AuthService{repository: repository}
}

func (s *AuthService) CreateEmployee(employee auth.Employee) (int, error) {
	employee.Password = generatePasswordHash(employee.Password)
	return s.repository.CreateEmployee(employee)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
