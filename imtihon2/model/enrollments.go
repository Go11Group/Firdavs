package model

import "time"

type Enrollments struct {
	EnrollmentId   string    `json:"enrollment_id`
	UserId         string    `json:"user_id"`
	CourseId       string    `json:"course_id`
	EnrollmentDate string    `json:"enrollment_date`
	CreatedAt      time.Time `json:"created_at"`
	UpdateAt       time.Time `json:"update_at`
	DeleteAt       int       `json:"delete_at`
}
