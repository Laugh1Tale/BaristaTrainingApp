package repository

import (
	"barTrApp/pkg/auth"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateEmployee(employee auth.Employee) (int, error)
	GetEmployee(email, password string) (auth.Employee, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
