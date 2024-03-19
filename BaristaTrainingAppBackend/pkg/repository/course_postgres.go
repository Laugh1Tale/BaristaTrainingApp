package repository

import (
	"github.com/jmoiron/sqlx"
)

type CoursePostgres struct {
	db *sqlx.DB
}

func NewCoursePostgres(db *sqlx.DB) *CoursePostgres {
	return &CoursePostgres{db: db}
}

// func (c *CoursePostgres) Create(course model.Course) (int, error) {
// }

// func (c *CoursePostgres) GetAll() ([]model.CourseResponse, error) {

// }

// func (c *CoursePostgres) GetById(courseId int) (model.CourseResponse, error) {

// }

func (c *CoursePostgres) Delete() {

}

func (c *CoursePostgres) Update() {

}
