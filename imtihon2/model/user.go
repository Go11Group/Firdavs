package model

import "time"

// User jadvalidagi foydalanuvchi ma'lumotlarni ifodalaydi
type User struct {
	Id        string    `json:"user_id"`    // Foydalanuvchi IDsi
	Name      string    `json:"name"`       // Ismi
	Email     string    `json:"email"`      // Elektron pochta manzili
	Birthday  string    `json:"birthday"`   // Tug'ilgan sanasi (string shaklida)
	Password  string    `json:"password"`   // Parol
	CreatedAt time.Time `json:"created_at"` // Yaratilgan vaqti
	UpdatedAt time.Time `json:"updated_at"` // Yangilangan vaqti
	DeleteAt  int       `json:"delete_at"`  // O'chirilgan vaqti (timestamp)
}

// FilterUsers jadvidagi foydalanuvchilarni filtrlash uchun ifodalaydi
type FilterUsers struct {
	Id       string `json:"user_id"`  // Foydalanuvchi IDsi
	Name     string `json:"name"`     // Ismi
	Email    string `json:"email"`    // Elektron pochta manzili
	Birthday string `json:"birthday"` // Tug'ilgan sanasi (string shaklida)
	Password string `json:"password"` // Parol
}

// FUsers jadvidagi foydalanuvchilarni qidirish uchun ifodalaydi
type FUsers struct {
	Id        string `json:"user_id"`    // Foydalanuvchi IDsi
	Name      string `json:"name"`       // Ismi
	Email     string `json:"email"`      // Elektron pochta manzili
	Birthday  string `json:"birthday"`   // Tug'ilgan sanasi (string shaklida)
	Password  string `json:"password"`   // Parol
	StartDate string `json:"start_date"` // Qidirishni boshlash vaqti
	EndDate   string `json:"end_date"`   // Qidirishni tugatish vaqti
}
