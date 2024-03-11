package repository

import (
	"barTrApp/pkg/model"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateEmployee(employee model.Employee) (int, error)
	GetEmployee(email, password string) (model.Employee, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
