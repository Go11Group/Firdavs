package postgres

import (
	"database/sql"
	"n11/Firdavs/imtihon2/model"

	"github.com/google/uuid"
)

// Courses struct o'z ichiga sql.DB obyektini saqlaydi.
type Courses struct {
	db *sql.DB
}

// NewCourses funktsiyasi yangi Courses obyekti qaytaradi.
func NewCourses(db *sql.DB) *Courses {
	return &Courses{db}
}

// CreateCourse funktsiyasi bazaga yangi kurs qo'shadigan metod.
func (u *Courses) CreateCourse(course *model.Courses) error {
	course.Course_id = uuid.NewString()
	_, err := u.db.Exec("insert into courses (course_id, title, description) VALUES ($1, $2, $3)",
		course.Course_id, course.Title, course.Description)
	if err != nil {
		return err
	}
	return nil
}

// ReadCourse funktsiyasi berilgan ID bo'yicha kursni o'qiydi.
func (u *Courses) ReadCourse(id string) (*model.Courses, error) {
	row := u.db.QueryRow("select * from courses where course_id = $1", id)
	var course model.Courses
	err := row.Scan(
		&course.Course_id, &course.Title, &course.Description, &course.CreatedAt, &course.UpdateAt, &course.DeleteAt)
	if err != nil {
		return nil, err
	}
	return &course, nil
}

// UpdateCourse funktsiyasi kursni yangilash uchun.
func (u *Courses) UpdateCourse(course *model.Courses) error {
	_, err := u.db.Exec("update courses set title = $1, description = $2 where course_id = $3",
		course.Title, course.Description, course.Course_id)
	if err != nil {
		return err
	}
	return nil
}

// DeleteCourse funktsiyasi kursni o'chirish uchun.
func (u *Courses) DeleteCourse(id string) error {
	_, err := u.db.Exec("delete from courses where course_id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

// ReadAllCourses funktsiyasi barcha kurslarni o'qish uchun.
func (u *Courses) ReadAllCourses() ([]*model.Courses, error) {
	rows, err := u.db.Query("select * from courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []*model.Courses
	for rows.Next() {
		var course model.Courses
		err := rows.Scan(
			&course.Course_id, &course.Title, &course.Description, &course.CreatedAt, &course.UpdateAt, &course.DeleteAt)
		if err != nil {
			return nil, err
		}
		courses = append(courses, &course)
	}
	return courses, nil
}

// FilterCourses funktsiyasi filtr bo'yicha kurslarni qidirish uchun.
func (u *Courses) FilterCourses(f model.FeltirCourses) ([]model.Courses, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
	)
	query := `select course_id, title, description from courses `
	filter := `where true`

	if len(f.Course_id) > 0 {
		params["course_id"] = f.Course_id
		filter += " and course_id = :course_id "
	}
	if len(f.Title) > 0 {
		params["title"] = f.Title
		filter += " and title = :title "
	}
	if len(f.Description) > 0 {
		params["description"] = f.Description
		filter += " and description = :description "
	}

	query = query + filter

	query, arr = ReplaceQueryParams(query, params)
	rows, err := u.db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []model.Courses
	for rows.Next() {
		var course model.Courses
		err = rows.Scan(&course.Course_id, &course.Title, &course.Description)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	return courses, nil
}
