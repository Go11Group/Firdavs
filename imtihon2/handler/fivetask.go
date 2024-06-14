package handler

import (
	"n11/Firdavs/imtihon2/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Task-1: GetCoursesByUser - foydalanuvchi ID'si bo'yicha kurslarni olish uchun handler funksiyasi
func (hand *Handler) GetCoursesByUser(g *gin.Context) {
	userID := g.Param("user_id")
	if userID == "" {
		// Agar foydalanuvchi ID'si kiritilmagan bo'lsa, 400 Bad Request status kodi bilan xatolikni qaytaradi
		g.String(http.StatusBadRequest, "User ID is required")
		return
	}

	// Foydalanuvchi ID'si bo'yicha kurslarni olish
	courses, err := hand.Fivetask.GetCoursesByUser(userID)
	if err != nil {
		// Agar xato bo'lsa, 500 Internal Server Error status kodi bilan xatolikni qaytaradi
		g.String(http.StatusInternalServerError, err.Error())
		return
	}

	// Kurslarni 200 OK status kodi bilan qaytaradi
	g.JSON(http.StatusOK, courses)
}

// Task-2: GetLessonsByCourse - kurs ID'si bo'yicha darslarni olish uchun handler funksiyasi
func (hand *Handler) GetLessonsByCourse(g *gin.Context) {
	courseID := g.Param("course_id")
	if courseID == "" {
		// Agar kurs ID'si kiritilmagan bo'lsa, 400 Bad Request status kodi bilan xatolikni qaytaradi
		g.String(http.StatusBadRequest, "Course ID is required")
		return
	}

	// Kurs ID'si bo'yicha darslarni olish
	lessons, err := hand.Fivetask.GetLessonsByCourse(courseID)
	if err != nil {
		// Agar xato bo'lsa, 500 Internal Server Error status kodi bilan xatolikni qaytaradi
		g.String(http.StatusInternalServerError, err.Error())
		return
	}

	// Darslarni 200 OK status kodi bilan qaytaradi
	g.JSON(http.StatusOK, lessons)
}

// Task-3: GetEnrolledUsersByCourse - kurs ID'si bo'yicha ro'yxatdan o'tgan foydalanuvchilarni olish uchun handler funksiyasi
func (hand *Handler) GetEnrolledUsersByCourse(g *gin.Context) {
	courseID := g.Param("course_id")
	if courseID == "" {
		// Agar kurs ID'si kiritilmagan bo'lsa, 400 Bad Request status kodi bilan xatolikni qaytaradi
		g.String(http.StatusBadRequest, "Course ID is required")
		return
	}

	// Kurs ID'si bo'yicha ro'yxatdan o'tgan foydalanuvchilarni olish
	users, err := hand.Fivetask.GetEnrolledUsersByCourse(courseID)
	if err != nil {
		// Agar xato bo'lsa, 500 Internal Server Error status kodi bilan xatolikni qaytaradi
		g.String(http.StatusInternalServerError, err.Error())
		return
	}

	// Foydalanuvchilarni 200 OK status kodi bilan qaytaradi
	g.JSON(http.StatusOK, users)
}

// Task-4: SearchUsers - foydalanuvchilarni filtr orqali qidirish uchun handler funksiyasi
func (hand *Handler) SearchUsers(g *gin.Context) {
	var userFilter model.FUsers
	// Query parametrlarini parslash
	if err := g.BindQuery(&userFilter); err != nil {
		// Agar xato bo'lsa, 400 Bad Request status kodi bilan xatolikni qaytaradi
		g.String(http.StatusBadRequest, err.Error())
		return
	}

	// Filtr orqali foydalanuvchilarni qidirish
	users, err := hand.Fivetask.SearchUsers(userFilter)
	if err != nil {
		// Agar xato bo'lsa, 500 Internal Server Error status kodi bilan xatolikni qaytaradi
		g.String(http.StatusInternalServerError, err.Error())
		return
	}

	// Foydalanuvchilarni 200 OK status kodi bilan qaytaradi
	g.JSON(http.StatusOK, users)
}
