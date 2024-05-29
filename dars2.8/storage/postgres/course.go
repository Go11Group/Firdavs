package postgres

import (
	"database/sql"
	"n11/Firdavs/dars2.8/model"
)

type CourseRepo struct {
	DB *sql.DB
}

func NewCourseRepo(DB *sql.DB) *CourseRepo {
	return &CourseRepo{DB}
}

func (c *CourseRepo) Create(course *model.Course) error {

	return nil
}
