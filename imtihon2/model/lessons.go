package model

import "time"

// Lessons jadvalidagi ma'lumotlarni ifodalaydi
type Lessons struct {
	LessonId  string    `json:"lesson_id"`  // Dars IDsi
	CourseId  string    `json:"course_id"`  // Kurs IDsi
	Title     string    `json:"title"`      // Sarlavha
	Content   string    `json:"content"`    // Tarkibi
	CreatedAt time.Time `json:"created_at"` // Yaratilgan vaqti
	UpdatedAt time.Time `json:"updated_at"` // Yangilangan vaqti
	DeleteAt  int       `json:"delete_at"`  // O'chirilgan vaqti (timestamp)
}

// FeltirLessons jadvidagi ma'lumotlarni filtrlash uchun ifodalaydi
type FeltirLessons struct {
	LessonId string `json:"lesson_id"` // Dars IDsi
	CourseId string `json:"course_id"` // Kurs IDsi
	Title    string `json:"title"`     // Sarlavha
	Content  string `json:"content"`   // Tarkibi
}
