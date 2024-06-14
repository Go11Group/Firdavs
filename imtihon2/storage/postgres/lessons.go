package postgres

import (
	"database/sql"
	"fmt"
	"n11/Firdavs/imtihon2/model"

	"github.com/google/uuid"
)

type Lesson struct {
	db *sql.DB
}

// NewLesson funktsiyasi yangi Lesson obyekti qaytaradi.
func NewLesson(db *sql.DB) *Lesson {
	return &Lesson{db}
}

// CreateLesson funktsiyasi yangi dars qo'shish uchun.
func (u *Lesson) CreateLesson(lesson *model.Lessons) error {
	lesson.LessonId = uuid.NewString()
	_, err := u.db.Exec("insert into lessons (lesson_id, course_id, title, content) VALUES ($1, $2, $3, $4)",
		lesson.LessonId, lesson.CourseId, lesson.Title, lesson.Content)
	if err != nil {
		return err
	}
	return nil
}

// ReadLesson funktsiyasi berilgan dars ID si bo'yicha ma'lumotni olish uchun.
func (u *Lesson) ReadLesson(id string) (*model.Lessons, error) {
	row := u.db.QueryRow("select * from lessons where lesson_id = $1", id)

	fmt.Println(id)
	var lesson model.Lessons
	err := row.Scan(
		&lesson.LessonId, &lesson.CourseId, &lesson.Title, &lesson.Content, &lesson.CreatedAt, &lesson.UpdatedAt, &lesson.DeleteAt)
	if err != nil {
		fmt.Println("58", err)
		return nil, err
	}
	return &lesson, nil
}

// UpdateLesson funktsiyasi darsni yangilash uchun.
func (u *Lesson) UpdateLesson(lesson *model.Lessons) error {
	_, err := u.db.Exec("update lessons set course_id = $1, title = $2, content = $3 where lesson_id = $4",
		lesson.CourseId, lesson.Title, lesson.Content, lesson.LessonId)
	if err != nil {
		return err
	}
	return nil
}

// DeleteLesson funktsiyasi darsni o'chirish uchun.
func (u *Lesson) DeleteLesson(id string) error {
	_, err := u.db.Exec("delete from lessons where lesson_id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

// GetAllLessons funktsiyasi barcha darslarni olish uchun.
func (u *Lesson) GetAllLessons() ([]*model.Lessons, error) {
	rows, err := u.db.Query("select * from lessons")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lessons []*model.Lessons
	for rows.Next() {
		var lesson model.Lessons
		err := rows.Scan(
			&lesson.LessonId, &lesson.CourseId, &lesson.Title, &lesson.Content, &lesson.CreatedAt, &lesson.UpdatedAt, &lesson.DeleteAt)
		if err != nil {
			return nil, err
		}
		lessons = append(lessons, &lesson)
	}
	return lessons, nil
}

// FilterLessons funktsiyasi darslarni filtrlash uchun.
func (u *Lesson) FilterLessons(f model.FeltirLessons) ([]model.Lessons, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
	)
	query := `select lesson_id, course_id, title, content
	from lessons `
	filter := ` where true`

	if len(f.LessonId) > 0 {
		params["lesson_id"] = f.LessonId
		filter += " and lesson_id = :lesson_id "
	}

	if len(f.CourseId) > 0 {
		params["course_id"] = f.CourseId
		filter += " and course_id = :course_id "
	}

	if len(f.Title) > 0 {
		params["title"] = f.Title
		filter += " and title = :title "
	}

	if len(f.Content) > 0 {
		params["content"] = f.Content
		filter += " and content = :content "
	}

	query = query + filter

	query, arr = ReplaceQueryParams(query, params)
	rows, err := u.db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lessons []model.Lessons
	for rows.Next() {
		var lesson model.Lessons
		err = rows.Scan(&lesson.LessonId, &lesson.CourseId, &lesson.Title, &lesson.Content)
		if err != nil {
			return nil, err
		}
		lessons = append(lessons, lesson)
	}
	return lessons, nil
}
