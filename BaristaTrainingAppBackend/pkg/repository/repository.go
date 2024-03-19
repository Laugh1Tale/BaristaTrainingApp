package repository

import (
	"barTrApp/pkg/model"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateEmployee(employee model.Employee) (int, error)
	GetEmployee(email, password string) (model.Employee, error)
}

type Course interface {
	Create()
	GetAll()
	GetById()
	Delete()
	Update()
}

type Lection interface {
	Create()
	GetAll()
	GetById()
	Delete()
	Update()
}

type Lesson interface {
	Create()
	GetAll()
	GetById()
	Delete()
	Update()
}

type Information interface {
	Create()
	GetAll()
	GetById()
	Delete()
	Update()
}

type Test interface {
	Create()
	GetAll()
	GetById()
	Delete()
	Update()
}

type Question interface {
	Create()
	GetAll()
	GetById()
	Delete()
	Update()
}

type Answer interface {
	Create()
	GetAll()
	GetById()
	Delete()
	Update()
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
