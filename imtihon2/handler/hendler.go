package handler

import (
	"database/sql"
	"n11/Firdavs/imtihon2/storage/postgres"
)

type Handler struct {
	db         *sql.DB
	User       *postgres.User
	Course     *postgres.Courses
	Lesson     *postgres.Lesson
	Enrollment *postgres.Enrollments
}

func NewHandler(db *sql.DB, user *postgres.User, course *postgres.Courses, lesson *postgres.Lesson, enrollments *postgres.Enrollments) *Handler {
	return &Handler{
		db:         db,
		User:       user,
		Course:     course,
		Lesson:     lesson,
		Enrollment: enrollments,
	}
}
