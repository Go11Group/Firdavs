package model

import "time"

type Lessons struct {
	LessonId  string    `json:"lesson_id`
	CourseId  string    `json:"course_id`
	Title     string    `json:"title`
	Content   string    `json:"content`
	CreatedAt time.Time `json:"created_at`
	UpdatedAt time.Time `json:"updated_at`
	DeleteAt  int       `json:"delete_at`
}
