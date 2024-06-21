package handler

import (
	"database/sql"
	"n11/Firdavs/imtihon2/storage/postgres"

	"github.com/gin-gonic/gin"
)

// GIN funksiyasi, Gin engine obyektini qaytaradi
func GIN(db *sql.DB, user *postgres.User, course *postgres.Courses, lesson *postgres.Lesson, enrollments *postgres.Enrollments, fivetask *postgres.Fivetask) *gin.Engine {
	r := gin.Default()

	handler := NewHandler(db, user, course, lesson, enrollments, fivetask)

	// User endpoints
	r.GET("/user", handler.GetAllUsers)             // Barcha foydalanuvchilarni olish
	r.GET("/user/:user_id", handler.GetUserById)    // ID bo'yicha foydalanuvchini olish
	r.GET("/userfilter", handler.GetFilterUsers)    // Foydalanuvchilarni filtr bo'yicha olish
	r.POST("/user", handler.CreateUser)             // Yangi foydalanuvchi yaratish
	r.PUT("/user/:user_id", handler.UpdateUser)     // Foydalanuvchini yangilash
	r.DELETE("/user/:user_id", handler.DeleteUsers) // ID bo'yicha foydalanuvchini o'chirish

	// Courses endpoints
	r.GET("/course", handler.GetAllCourses)              // Barcha kurslarni olish
	r.GET("/course/:course_id", handler.GetCourseById)   // ID bo'yicha kursni olish
	r.GET("/coursefilter", handler.GetFilterCourses)     // Kurslarni filtr bo'yicha olish
	r.POST("/course", handler.CreateCourse)              // Yangi kurs yaratish
	r.PUT("/course/", handler.UpdateCourse)              // Kursni yangilash
	r.DELETE("/course/:course_id", handler.DeleteCourse) // ID bo'yicha kursni o'chirish

	// Lesson endpoints
	r.GET("/lesson", handler.GetAllLessons)              // Barcha darslarni olish
	r.GET("/lesson/:lesson_id", handler.GetLessonById)   // ID bo'yicha darsni olish
	r.POST("/lesson", handler.CreateLesson)              // Yangi dars yaratish
	r.GET("/lesson/filter", handler.GetFilterLessons)    // Darslarni filtr bo'yicha olish
	r.PUT("/lesson/", handler.UpdateLesson)              // Darsni yangilash
	r.DELETE("/lesson/:lesson_id", handler.DeleteLesson) // ID bo'yicha darsni o'chirish

	// Enrollments endpoints
	r.GET("/enrollments", handler.GetAllEnrollments)                  // Barcha enrollments (ro'yxatdan o'tishlar)ni olish
	r.GET("/enrollments/:enrollment_id", handler.GetEnrollmentById)   // ID bo'yicha enrollmentni olish
	r.POST("/enrollments", handler.CreateEnrollments)                 // Yangi enrollment yaratish
	r.GET("/enrollments/filter", handler.GetFilterEnrollments)        // Enrollmentslarni filtr bo'yicha olish
	r.PUT("/enrollments/", handler.UpdateEnrollment)                  // Enrollmentni yangilash
	r.DELETE("/enrollments/:enrollment_id", handler.DeleteEnrollment) // ID bo'yicha enrollmentni o'chirish

	// Additional endpoints
	r.GET("/courses/user/:user_id", handler.GetCoursesByUser)           // Foydalanuvchi bo'yicha kurslarni olish
	r.GET("/lessons/course/:course_id", handler.GetLessonsByCourse)     // Kurs bo'yicha darslarni olish
	r.GET("/users/course/:course_id", handler.GetEnrolledUsersByCourse) // Kurs bo'yicha enrollments foydalanuvchilarni olish
	r.GET("/search/users", handler.SearchUsers)                         // Foydalanuvchilarni qidirish

	return r
}
