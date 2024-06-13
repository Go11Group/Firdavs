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

func NewEnrollments(db *sql.DB) *Enrollments {
	return &Enrollments{db}
}

func (u *Enrollments) CreateEnrollments(enrrolment *model.Enrollments) error {

	enrollmentDate := time.Now()


	enrrolment.EnrollmentId = uuid.NewString()
	_, err := u.db.Exec("insert into enrollments (enrollment_id, user_id, course_id, enrolment_date) VALUES ($1,$2,$3,$4)",
		enrrolment.EnrollmentId, enrrolment.UserId, enrrolment.CourseId, enrollmentDate)
	if err != nil {

		fmt.Println("3hmmmmnn2", err)
		return err
	}
	return nil
}



func (u *Enrollments) GetEnrollment(id string) (*model.Enrollments, error) {
	row := u.db.QueryRow("select * from enrollments where enrollment_id = $1", id)

	fmt.Println(id)
	var enrrolment model.Enrollments
	err := row.Scan(
		&enrrolment.EnrollmentId, &enrrolment.UserId, &enrrolment.CourseId, &enrrolment.EnrollmentDate,  &enrrolment.CreatedAt, &enrrolment.UpdateAt, &enrrolment.DeleteAt)
	if err != nil {
		fmt.Println("64", err)
		return nil, err
	}
	return &enrrolment, nil
}

func (u *Enrollments) UpdateEnrollment(enrrolment *model.Enrollments) error {
	enrollmentDate, err := time.Parse("2006-01-02", enrrolment.EnrollmentDate)
	if err != nil {
		return errors.New("failed to parse enrollment date")
	}
	_, err = u.db.Exec("update enrollments set user_id = $1, course_id = $2, enrollment_date = $3 where enrollment_id = $5",
		enrrolment.UserId, enrrolment.CourseId, enrollmentDate, enrrolment.EnrollmentId)
	if err != nil {
		return err
	}
	return nil
}

func (u *Enrollments) DeleteEnrollment(id string) error {
	_, err := u.db.Exec("delete from enrollments where enrollment_id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (u *Enrollments) ReadAllEnrollments() ([]*model.Enrollments, error) {
	rows, err := u.db.Query("select * from enrollments")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var enrollments []*model.Enrollments
	for rows.Next() {
		var enrrolment model.Enrollments
		err := rows.Scan(
			&enrrolment.EnrollmentId, &enrrolment.UserId, &enrrolment.CourseId, &enrrolment.EnrollmentDate,  &enrrolment.CreatedAt, &enrrolment.UpdateAt, &enrrolment.DeleteAt)
		if err != nil {
			return nil, err
		}
		enrollments = append(enrollments, &enrrolment)
	}
	return enrollments, nil
}

func (u *Enrollments) FilterUsers(f model.FeltirEnrollments) ([]model.Enrollments, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
	)
	query := `select enrollment_id, user_id, course_id, enrollment_date
  from enrollments `
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
		return nil, err
	}
	defer rows.Close()

	var enrollments []model.Enrollments
	for rows.Next() {
		var enrrolment model.Enrollments
		err = rows.Scan(&enrrolment.EnrollmentId, &enrrolment.UserId, &enrrolment.CourseId, &enrrolment.EnrollmentDate)
		if err != nil {
			return nil, err
		}

		enrollments = append(enrollments, enrrolment)
	}
	return enrollments, nil
}
