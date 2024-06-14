package model

import "time"

// Courses - Kurslar ma'lumotlarini saqlash uchun struktura
type Courses struct {
	Course_id   string    `json:"course_id"`   // Kurs ID-si
	Title       string    `json:"title"`       // Kurs nomi
	Description string    `json:"description"` // Kurs ta'rifi
	CreatedAt   time.Time `json:"created_at"`  // Yaratilgan sana
	UpdateAt    time.Time `json:"update_at"`   // Yangilangan sana
	DeleteAt    int       `json:"delete_at"`   // O'chirilgan sana (0 bo'lsa, kurs yo'qolmagan)
}

// FeltirCourses - Kurslarni filtrlash uchun foydalaniladigan struktura
type FeltirCourses struct {
	Course_id   string // Kurs ID-si
	Title       string // Kurs nomi
	Description string // Kurs ta'rifi
}
