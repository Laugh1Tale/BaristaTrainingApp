package repository

import (
	"barTrApp/pkg/auth"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateEmployee(employee auth.Employee) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, lastname, email, password_hash) values ($1, $2, $3, $4) RETURNING id", employeeTable)
	row := r.db.QueryRow(query, employee.Name, employee.LastName, employee.Email, employee.Password)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}