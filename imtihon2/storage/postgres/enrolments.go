package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"n11/Firdavs/imtihon2/model"
	"time"

	"github.com/google/uuid"
)

type Enrollments struct {
	db *sql.DB
}

// NewEnrollments funktsiyasi yangi Enrollments obyekti qaytaradi.
func NewEnrollments(db *sql.DB) *Enrollments {
	return &Enrollments{db}
}

// CreateEnrollment funktsiyasi bazaga yangi yozuv qo'shadigan metod.
func (u *Enrollments) CreateEnrollment(enrollment *model.Enrollments) error {
	enrollment.EnrollmentId = uuid.NewString()
	enrollment.EnrollmentDate = time.Now().Format("2006-01-02")
	_, err := u.db.Exec("insert into enrollments (enrollment_id, user_id, course_id, enrollment_date) VALUES ($1, $2, $3, $4)",
		enrollment.EnrollmentId, enrollment.UserId, enrollment.CourseId, enrollment.EnrollmentDate)
	if err != nil {
		fmt.Println("Error while creating enrollment:", err)
		return err
	}
	return nil
}

// GetEnrollment funktsiyasi berilgan ID bo'yicha yozuvni o'qish uchun.
func (u *Enrollments) GetEnrollment(id string) (*model.Enrollments, error) {
	row := u.db.QueryRow("select * from enrollments where enrollment_id = $1", id)

	fmt.Println("Fetching enrollment with ID:", id)
	var enrollment model.Enrollments
	err := row.Scan(
		&enrollment.EnrollmentId, &enrollment.UserId, &enrollment.CourseId, &enrollment.EnrollmentDate, &enrollment.CreatedAt, &enrollment.UpdateAt, &enrollment.DeleteAt)
	if err != nil {
		fmt.Println("Error while fetching enrollment:", err)
		return nil, err
	}
	return &enrollment, nil
}

// UpdateEnrollment funktsiyasi yozuvni yangilash uchun.
func (u *Enrollments) UpdateEnrollment(enrollment *model.Enrollments) error {
	enrollmentDate, err := time.Parse("2006-01-02", enrollment.EnrollmentDate)
	if err != nil {
		return errors.New("failed to parse enrollment date")
	}
	_, err = u.db.Exec("update enrollments set user_id = $1, course_id = $2, enrollment_date = $3 where enrollment_id = $4",
		enrollment.UserId, enrollment.CourseId, enrollmentDate, enrollment.EnrollmentId)
	if err != nil {
		fmt.Println("Error while updating enrollment:", err)
		return err
	}
	return nil
}

// DeleteEnrollment funktsiyasi yozuvni o'chirish uchun.
func (u *Enrollments) DeleteEnrollment(id string) error {
	_, err := u.db.Exec("delete from enrollments where enrollment_id = $1", id)
	if err != nil {
		fmt.Println("Error while deleting enrollment:", err)
		return err
	}
	return nil
}

// ReadAllEnrollments funktsiyasi barcha yozuvlarni o'qish uchun.
func (u *Enrollments) ReadAllEnrollments() ([]*model.Enrollments, error) {
	rows, err := u.db.Query("select * from enrollments")
	if err != nil {
		fmt.Println("Error while fetching all enrollments:", err)
		return nil, err
	}
	defer rows.Close()

	var enrollments []*model.Enrollments
	for rows.Next() {
		var enrollment model.Enrollments
		err := rows.Scan(
			&enrollment.EnrollmentId, &enrollment.UserId, &enrollment.CourseId, &enrollment.EnrollmentDate, &enrollment.CreatedAt, &enrollment.UpdateAt, &enrollment.DeleteAt)
		if err != nil {
			fmt.Println("Error while scanning enrollment row:", err)
			return nil, err
		}
		enrollments = append(enrollments, &enrollment)
	}
	return enrollments, nil
}

// FilterEnrollment funktsiyasi filtr bo'yicha yozuvlarni qidirish uchun.
func (u *Enrollments) FilterEnrollment(f model.FeltirEnrollments) ([]model.Enrollments, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
	)
	query := `select enrollment_id, user_id, course_id, enrollment_date from enrollments `
	filter := ` where true`

	if len(f.EnrollmentId) > 0 {
		params["enrollment_id"] = f.EnrollmentId
		filter += " and enrollment_id = :enrollment_id "
	}

	if len(f.UserId) > 0 {
		params["user_id"] = f.UserId
		filter += " and user_id = :user_id "
	}

	if len(f.CourseId) > 0 {
		params["course_id"] = f.CourseId
		filter += " and course_id = :course_id "
	}

	if len(f.EnrollmentDate) > 0 {
		params["enrollment_date"] = f.EnrollmentDate
		filter += " and enrollment_date = :enrollment_date "
	}

	query = query + filter

	query, arr = ReplaceQueryParams(query, params)
	rows, err := u.db.Query(query, arr...)
	if err != nil {
		fmt.Println("Error while filtering enrollments:", err)
		return nil, err
	}
	defer rows.Close()

	var enrollments []model.Enrollments
	for rows.Next() {
		var enrollment model.Enrollments
		err = rows.Scan(&enrollment.EnrollmentId, &enrollment.UserId, &enrollment.CourseId, &enrollment.EnrollmentDate)
		if err != nil {
			fmt.Println("Error while scanning filtered enrollment row:", err)
			return nil, err
		}

		enrollments = append(enrollments, enrollment)
	}
	return enrollments, nil
}
