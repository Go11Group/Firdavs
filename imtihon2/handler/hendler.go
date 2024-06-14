package handler

import (
	"database/sql"
	"n11/Firdavs/imtihon2/storage/postgres"
)

// Handler strukturasi barcha zarur bo'lgan komponentlarni o'z ichiga oladi
type Handler struct {
	db         *sql.DB               // SQL ma'lumotlar bazasi uchun ulanish
	User       *postgres.User        // User modeli uchun handler
	Course     *postgres.Courses     // Course modeli uchun handler
	Lesson     *postgres.Lesson      // Lesson modeli uchun handler
	Enrollment *postgres.Enrollments // Enrollments modeli uchun handler
	Fivetask   *postgres.Fivetask    // Fivetask modeli uchun handler
}

// NewHandler funksiyasi yangi Handler obyektini yaratadi
func NewHandler(db *sql.DB, user *postgres.User, course *postgres.Courses, lesson *postgres.Lesson, enrollments *postgres.Enrollments, fivetask *postgres.Fivetask) *Handler {
	return &Handler{
		db:         db,          // SQL ma'lumotlar bazasi ulanishini o'rnatadi
		User:       user,        // User modeli handlerini o'rnatadi
		Course:     course,      // Course modeli handlerini o'rnatadi
		Lesson:     lesson,      // Lesson modeli handlerini o'rnatadi
		Enrollment: enrollments, // Enrollments modeli handlerini o'rnatadi
		Fivetask:   fivetask,    // Fivetask modeli handlerini o'rnatadi
	}
}
