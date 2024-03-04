package repository

import (
	"barTrApp/pkg/auth"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateEmployee(employee auth.Employee) (int, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
