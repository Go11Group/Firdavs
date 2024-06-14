package model

import "time"

// Enrollments enrollments jadvalidagi ma'lumotlarni ifodalaydi
type Enrollments struct {
	EnrollmentId   string    `json:"enrollment_id"`   // Enrollments ID
	UserId         string    `json:"user_id"`         // Foydalanuvchi IDsi
	CourseId       string    `json:"course_id"`       // Kurs IDsi
	EnrollmentDate string    `json:"enrollment_date"` // Ro'yxatdan o'tish sanasi
	CreatedAt      time.Time `json:"created_at"`      // Yaratilgan vaqti
	UpdateAt       time.Time `json:"update_at"`       // Yangilangan vaqti
	DeleteAt       int       `json:"delete_at"`       // O'chirilgan vaqti (timestamp)
}

// FeltirEnrollments enrollments jadvidagi ma'lumotlarni filtrlash uchun ifodalaydi
type FeltirEnrollments struct {
	EnrollmentId   string `json:"enrollment_id"`   // Enrollments ID
	UserId         string `json:"user_id"`         // Foydalanuvchi IDsi
	CourseId       string `json:"course_id"`       // Kurs IDsi
	EnrollmentDate string `json:"enrollment_date"` // Ro'yxatdan o'tish sanasi
}
