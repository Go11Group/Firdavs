package model

import "time"

type Courses struct {
	Course_id   string    `json:"course_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at`
	UpdateAt    time.Time `json:"update_at"`
	DeleteAt    int       `json:"delete_at"`
}
