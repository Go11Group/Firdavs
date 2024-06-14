package handler

import (
	"n11/Firdavs/imtihon2/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateEnrollments yangi enrollments (ro'yxatdan o'tish) yaratish uchun ishlatiladi
func (hand *Handler) CreateEnrollments(g *gin.Context) {
	var enrollment model.Enrollments
	// JSON ma'lumotlarini model.Enrollments strukturasiga yuklaydi
	err := g.BindJSON(&enrollment)
	if err != nil {
		// Agar xato yuz bersa, 400 Bad Request status kodi bilan xatolikni qaytaradi
		g.String(http.StatusBadRequest, err.Error())
		return
	}
	// Enrollmentsni bazaga qo'shish
	err = hand.Enrollment.CreateEnrollment(&enrollment)
	if err != nil {
		// Agar xato yuz bersa, 500 Internal Server Error status kodi bilan xatolikni qaytaradi
		g.String(http.StatusInternalServerError, err.Error())
		return
	}
	// Enrollments yaratildi va 201 Created status kodi bilan qaytariladi
	g.JSON(http.StatusCreated, enrollment)
}

// UpdateEnrollment mavjud enrollmentsni yangilash uchun ishlatiladi
func (hand *Handler) UpdateEnrollment(g *gin.Context) {
	var enrollment model.Enrollments
	// JSON ma'lumotlarini model.Enrollments strukturasiga yuklaydi
	err := g.Bind(&enrollment)
	if err != nil {
		// Agar xato yuz bersa, 400 Bad Request status kodi bilan xatolikni qaytaradi
		g.String(http.StatusBadRequest, err.Error())
		return
	}
	// Enrollmentsni yangilash
	err = hand.Enrollment.UpdateEnrollment(&enrollment)
	if err != nil {
		// Agar xato yuz bersa, 500 Internal Server Error status kodi bilan xatolikni qaytaradi
		g.String(http.StatusInternalServerError, err.Error())
		return
	}
	// Enrollments muvaffaqiyatli yangilandi va 200 OK status kodi bilan xabar qaytariladi
	g.String(200, "SUCCESS")
}

// GetAllEnrollments barcha enrollmentslarni olish uchun ishlatiladi
func (hand *Handler) GetAllEnrollments(g *gin.Context) {
	// Barcha enrollmentslarni bazadan o'qish
	enrollments, err := hand.Enrollment.ReadAllEnrollments()
	if err != nil {
		// Agar xato yuz bersa, 500 Internal Server Error status kodi bilan xatolikni qaytaradi
		g.String(http.StatusInternalServerError, err.Error())
		return
	}
	// Enrollmentslarni JSON formatida 200 OK status kodi bilan qaytarish
	g.JSON(http.StatusOK, enrollments)
}

// GetEnrollmentById enrollmentsni ID bo'yicha olish uchun ishlatiladi
func (hand *Handler) GetEnrollmentById(g *gin.Context) {
	id := g.Param("enrollment_id")
	if id == "" {
		// Agar ID bo'sh bo'lsa, 400 Bad Request status kodi bilan xatolikni qaytaradi
		g.String(http.StatusBadRequest, "There is an error entering ID")
		return
	}
	// Enrollmentsni bazadan o'qish
	enrollment, err := hand.Enrollment.GetEnrollment(id)
	if err != nil {
		// Agar xato yuz bersa, 500 Internal Server Error status kodi bilan xatolikni qaytaradi
		g.String(http.StatusInternalServerError, err.Error())
		return
	}
	// Enrollmentsni JSON formatida 200 OK status kodi bilan qaytarish
	g.JSON(http.StatusOK, enrollment)
}

// DeleteEnrollment enrollmentsni ID bo'yicha o'chirish uchun ishlatiladi
func (hand *Handler) DeleteEnrollment(g *gin.Context) {
	id := g.Param("enrollment_id")
	if id == "" {
		// Agar ID bo'sh bo'lsa, 400 Bad Request status kodi bilan xatolikni qaytaradi
		g.String(http.StatusBadRequest, "There is an error entering ID")
		return
	}
	// Enrollmentsni bazadan o'chirish
	err := hand.Enrollment.DeleteEnrollment(id)
	if err != nil {
		// Agar xato yuz bersa, 500 Internal Server Error status kodi bilan xatolikni qaytaradi
		g.String(http.StatusInternalServerError, err.Error())
		return
	}
	// Enrollments muvaffaqiyatli o'chirildi va 200 OK status kodi bilan xabar qaytariladi
	g.String(200, "SUCCESS")
}

// GetFilterEnrollments enrollmentslarni filtrlash uchun ishlatiladi
func (hand *Handler) GetFilterEnrollments(g *gin.Context) {
	enrollmentFilter := model.FeltirEnrollments{}
	// Filtr parametrlari olish
	enrollmentFilter.EnrollmentId = g.Param("enrollment_id")
	enrollmentFilter.UserId = g.Param("user_id")
	enrollmentFilter.CourseId = g.Param("course_id")
	enrollmentFilter.EnrollmentDate = g.Param("enrollment_date")

	// Enrollmentslarni filtrlash
	enrollment, err := hand.Enrollment.FilterEnrollment(enrollmentFilter)
	if err != nil {
		// Agar xato yuz bersa, 400 Bad Request status kodi bilan xatolikni qaytaradi
		g.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// Filtrlangan enrollmentslarni JSON formatida 200 OK status kodi bilan qaytarish
	g.JSON(http.StatusOK, enrollment)
}
