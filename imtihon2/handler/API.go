package handler

import (
	"database/sql"
	"n11/Firdavs/imtihon2/storage/postgres"

	"github.com/gin-gonic/gin"
)

func GIN(db *sql.DB, user *postgres.User, course *postgres.Courses, lesson *postgres.Lesson, enrollments *postgres.Enrollments) *gin.Engine {
	r := gin.Default()

	handler := NewHandler(db, user, course, lesson, enrollments)

	r.GET("/user", handler.GetAllUsers)
	r.GET("/user/:user_id", handler.GetUserById)
	r.GET("/user/filter", handler.GetFilterUsers)
	r.POST("/user", handler.CreateUser)
	r.PUT("/user", handler.UpdateUser)
	r.DELETE("/user/:user_id", handler.DeleteUsers)

	r.GET("/course", handler.GetAllCourses)
	r.GET("/course/:course_id", handler.GetCourseById)
	r.GET("/coursefilter", handler.GetFilterCourses)
	r.POST("/course", handler.CreateCurse)
	r.PUT("/course/", handler.UpdateCourse)
	r.DELETE("/course/:id", handler.DeleteCourse)

	r.GET("/lesson", handler.GetAllLessons)
	r.GET("/lesson/:lesson_id", handler.GetLessonById)
	r.POST("/lesson", handler.CreateLesson)
	r.GET("/lesson/filter", handler.GetFilterLessons)
	r.PUT("/lesson/", handler.UpdateLesson)
	r.DELETE("/lesson/:id", handler.DeleteLesson)

	r.GET("/enrollments", handler.GetAllEnrollments)
	r.GET("/enrollments/:enrollments_id", handler.GetEnrollmentById)
	r.POST("/enrollments", handler.CreateEnrollments)
	r.GET("/enrollments/filter", handler.GetFilterEnrollments)
	r.PUT("/enrollments/", handler.UpdateEnrollment)
	r.DELETE("/enrollments/:id", handler.DeleteEnrollment)

	return r
}
