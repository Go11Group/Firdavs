package postgres

import (
	"database/sql"
	"n11/Firdavs/imtihon2/model"
)

type Fivetask struct {
	db *sql.DB
}

// NewFivetask funktsiyasi yangi Fivetask obyekti qaytaradi.
func NewFivetask(db *sql.DB) *Fivetask {
	return &Fivetask{db}
}

// * Task-1
// GetCoursesByUser funktsiyasi berilgan foydalanuvchi ID si bo'yicha kurslarni olish uchun.
func (u *Fivetask) GetCoursesByUser(userID string) ([]model.Courses, error) {
	query := `
	SELECT c.course_id, c.title, c.description, c.created_at, c.updated_at 
	FROM courses c 
	JOIN enrollments e ON c.course_id = e.course_id 
	WHERE e.user_id = $1`

	rows, err := u.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []model.Courses
	for rows.Next() {
		var course model.Courses
		err = rows.Scan(&course.Course_id, &course.Title, &course.Description, &course.CreatedAt, &course.UpdateAt)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return courses, nil
}

// * Task-2
// GetLessonsByCourse funktsiyasi berilgan kurs ID si bo'yicha darslarni olish uchun.
func (u *Fivetask) GetLessonsByCourse(courseID string) ([]model.Lessons, error) {
	query := `
	SELECT l.lesson_id, l.title, l.content, l.created_at, l.updated_at 
	FROM lessons l 
	WHERE l.course_id = $1`

	rows, err := u.db.Query(query, courseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lessons []model.Lessons
	for rows.Next() {
		var lesson model.Lessons
		err = rows.Scan(&lesson.LessonId, &lesson.Title, &lesson.Content, &lesson.CreatedAt, &lesson.UpdatedAt)
		if err != nil {
			return nil, err
		}
		lessons = append(lessons, lesson)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return lessons, nil
}

// * Task-3
// GetEnrolledUsersByCourse funktsiyasi berilgan kurs ID si bo'yicha ro'yxatdan o'tgan foydalanuvchilarni olish uchun.
func (u *Fivetask) GetEnrolledUsersByCourse(courseID string) ([]model.User, error) {
	query := `
	SELECT u.user_id, u.name, u.email, u.birthday, u.password, u.created_at 
	FROM users u 
	JOIN enrollments e ON u.user_id = e.user_id 
	WHERE e.course_id = $1`

	rows, err := u.db.Query(query, courseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Birthday, &user.Password, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// * Task-4
// SearchUsers funktsiyasi foydalanuvchilarni filtrlash uchun.
func (u *Fivetask) SearchUsers(f model.FUsers) ([]model.User, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
	)
	query := `SELECT user_id, name, email, birthday, password FROM users`
	filter := ` WHERE true`

	if len(f.Id) > 0 {
		params["user_id"] = f.Id
		filter += " AND user_id = :user_id"
	}
	if len(f.Name) > 0 {
		params["name"] = f.Name
		filter += " AND name ILIKE :name"
	}
	if len(f.Email) > 0 {
		params["email"] = f.Email
		filter += " AND email ILIKE :email"
	}
	if len(f.Birthday) > 0 {
		params["birthday"] = f.Birthday
		filter += " AND birthday = :birthday"
	}
	if len(f.StartDate) > 0 && len(f.EndDate) > 0 {
		params["start_date"] = f.StartDate
		params["end_date"] = f.EndDate
		filter += " AND birthday BETWEEN :start_date AND :end_date"
	}

	query = query + filter

	query, arr = ReplaceQueryParams(query, params)
	rows, err := u.db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Birthday, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
