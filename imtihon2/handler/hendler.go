package handler

import (
	"database/sql"
	"n11/Firdavs/imtihon2/storage/postgres"
)

type Handler struct {
	db         *sql.DB
	User       *postgres.User
	Course     *postgres.Courses
	// Lesson     *postgres.Lessons
	// Enrollment *postgres.Enrollments
}

func NewHandler(db *sql.DB, user *postgres.User, course *postgres.Courses) *Handler {
	return &Handler{
		db:         db,
		User:       user,
		Course:     course,
		// Lesson:     lesson,
		// Enrollment: enrollment,
	}
}
