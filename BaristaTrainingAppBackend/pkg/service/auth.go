package service

import (
	"barTrApp/pkg/auth"
	"barTrApp/pkg/repository"
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	salt       = "kahj7832jkasdfh23"
	signingKey = "asdgh34hjk289asdg23ds215"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	EmployeeId int `json:"employee_id"`
}

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

func (s *AuthService) GenerateToken(email, password string) (string, error) {
	employee, err := s.repository.GetEmployee(email, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		employee.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
