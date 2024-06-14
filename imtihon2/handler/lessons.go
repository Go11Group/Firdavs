package handler

import (
	"n11/Firdavs/imtihon2/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateLesson - yangi dars yaratish uchun handler funksiyasi
func (hand *Handler) CreateLesson(g *gin.Context) {
	var lesson model.Lessons
	// JSON requestini parslash
	err := g.BindJSON(&lesson)
	if err != nil {
		// Agar xato bo'lsa, 400 Bad Request status kodi bilan xatolikni qaytaradi
		g.String(http.StatusBadRequest, err.Error())
		return
	}
	// Darsni yaratish
	err = hand.Lesson.CreateLesson(&lesson)
	if err != nil {
		// Agar xato bo'lsa, 500 Internal Server Error status kodi bilan xatolikni qaytaradi
		g.String(http.StatusInternalServerError, err.Error())
		return
	}
	// Yaratilgan darsni 201 Created status kodi bilan qaytaradi
	g.JSON(http.StatusCreated, lesson)
}

// UpdateLesson - mavjud darsni yangilash uchun handler funksiyasi
func (hand *Handler) UpdateLesson(g *gin.Context) {
	var lesson model.Lessons
	// Requestni parslash
	err := g.Bind(&lesson)
	if err != nil {
		// Agar xato bo'lsa, 400 Bad Request status kodi bilan xatolikni qaytaradi
		g.String(http.StatusBadRequest, err.Error())
		return
	}
	// Darsni yangilash
	err = hand.Lesson.UpdateLesson(&lesson)
	if err != nil {
		// Agar xato bo'lsa, 500 Internal Server Error status kodi bilan xatolikni qaytaradi
		g.String(http.StatusInternalServerError, err.Error())
		return
	}
	// Muvaffaqiyatli yangilanganini bildiradi
	g.String(http.StatusOK, "SUCCESS")
}

// GetAllLessons - barcha darslarni olish uchun handler funksiyasi
func (hand *Handler) GetAllLessons(g *gin.Context) {
	// Barcha darslarni olish
	lessons, err := hand.Lesson.GetAllLessons()
	if err != nil {
		// Agar xato bo'lsa, 500 Internal Server Error status kodi bilan xatolikni qaytaradi
		g.String(http.StatusInternalServerError, err.Error())
		return
	}
	// Barcha darslarni 200 OK status kodi bilan qaytaradi
	g.JSON(http.StatusOK, lessons)
}

// GetLessonById - darsni ID orqali olish uchun handler funksiyasi
func (hand *Handler) GetLessonById(g *gin.Context) {
	id := g.Param("lesson_id")
	if id == "" {
		// Agar ID kiritilmagan bo'lsa, 400 Bad Request status kodi bilan xatolikni qaytaradi
		g.String(http.StatusBadRequest, "There is an error entering ID")
		return
	}

	// Darsni ID orqali olish
	lesson, err := hand.Lesson.ReadLesson(id)
	if err != nil {
		// Agar xato bo'lsa, 500 Internal Server Error status kodi bilan xatolikni qaytaradi
		g.String(http.StatusInternalServerError, err.Error())
		return
	}
	// Darsni 200 OK status kodi bilan qaytaradi
	g.JSON(http.StatusOK, lesson)
}

// DeleteLesson - darsni ID orqali o'chirish uchun handler funksiyasi
func (hand *Handler) DeleteLesson(g *gin.Context) {
	id := g.Param("lesson_id")
	if id == "" {
		// Agar ID kiritilmagan bo'lsa, 400 Bad Request status kodi bilan xatolikni qaytaradi
		g.String(http.StatusBadRequest, "There is an error entering ID")
		return
	}

	// Darsni o'chirish
	err := hand.Lesson.DeleteLesson(id)
	if err != nil {
		// Agar xato bo'lsa, 500 Internal Server Error status kodi bilan xatolikni qaytaradi
		g.String(http.StatusInternalServerError, err.Error())
		return
	}
	// Muvaffaqiyatli o'chirilganini bildiradi
	g.String(http.StatusOK, "SUCCESS")
}

// GetFilterLessons - darslarni filtrlar orqali olish uchun handler funksiyasi
func (hand *Handler) GetFilterLessons(g *gin.Context) {
	// Filtr parametrlari
	LessonFilter := model.FeltirLessons{
		LessonId: g.Param("lesson_id"),
		CourseId: g.Param("course_id"),
		Title:    g.Param("title"),
		Content:  g.Param("content"),
	}

	// Filtrlangan darslarni olish
	lesson, err := hand.Lesson.FilterLessons(LessonFilter)
	if err != nil {
		// Agar xato bo'lsa, 400 Bad Request status kodi bilan xatolikni qaytaradi
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Filtrlangan darslarni 200 OK status kodi bilan qaytaradi
	g.JSON(http.StatusOK, lesson)
}
