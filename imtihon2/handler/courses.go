package handler

import (
	"fmt"
	"n11/Firdavs/imtihon2/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateCourse yangi kurs yaratish uchun ishlatiladi
func (hand *Handler) CreateCourse(g *gin.Context) {
	var course model.Courses
	// JSON ma'lumotlarini model.Courses strukturasiga yuklaydi
	err := g.BindJSON(&course)
	if err != nil {
		g.String(http.StatusBadRequest, err.Error())
		return
	}
	// Kursni bazaga qo'shish
	err = hand.Course.CreateCourse(&course)
	if err != nil {
		g.String(http.StatusInternalServerError, err.Error())
		return
	}
	// Kurs yaratildi
	g.JSON(http.StatusCreated, course)
}

// UpdateCourse mavjud kursni yangilash uchun ishlatiladi
func (hand *Handler) UpdateCourse(g *gin.Context) {
	var course model.Courses
	// JSON ma'lumotlarini model.Courses strukturasiga yuklaydi
	err := g.Bind(&course)
	if err != nil {
		g.String(http.StatusBadRequest, err.Error())
		return
	}
	// Kursni yangilash
	err = hand.Course.UpdateCourse(&course)
	if err != nil {
		g.String(http.StatusInternalServerError, err.Error())
		return
	}
	// Kurs muvaffaqiyatli yangilandi
	g.String(200, "SUCCESS")
}

// GetAllCourses barcha kurslarni olish uchun ishlatiladi
func (hand *Handler) GetAllCourses(g *gin.Context) {
	// Barcha kurslarni bazadan o'qish
	courses, err := hand.Course.ReadAllCourses()
	if err != nil {
		g.String(http.StatusInternalServerError, err.Error())
		return
	}
	// Kurslarni JSON formatida qaytarish
	g.JSON(http.StatusOK, courses)
}

// GetCourseById kursni ID bo'yicha olish uchun ishlatiladi
func (hand *Handler) GetCourseById(g *gin.Context) {
	id := g.Param("course_id")
	if id == "" {
		g.String(http.StatusBadRequest, "There is an error entering ID")
		return
	}
	// Kursni bazadan o'qish
	course, err := hand.Course.ReadCourse(id)
	if err != nil {
		fmt.Println("64 hand", err)
		g.String(http.StatusInternalServerError, err.Error())
		return
	}
	// Kursni JSON formatida qaytarish
	g.JSON(http.StatusOK, course)
}

// DeleteCourse kursni ID bo'yicha o'chirish uchun ishlatiladi
func (hand *Handler) DeleteCourse(g *gin.Context) {
	id := g.Param("course_id")
	if id == "" {
		g.String(http.StatusBadRequest, "There is an error entering ID")
		return
	}
	// Kursni bazadan o'chirish
	err := hand.Course.DeleteCourse(id)
	if err != nil {
		g.String(http.StatusInternalServerError, err.Error())
		return
	}
	// Kurs muvaffaqiyatli o'chirildi
	g.String(200, "SUCCESS")
}

// GetFilterCourses kurslarni filtrlash uchun ishlatiladi
func (hand *Handler) GetFilterCourses(g *gin.Context) {
	courseFilter := model.FeltirCourses{}
	// Filtr parametrlari
	courseFilter.Course_id = g.Param("course_id")
	courseFilter.Title = g.Param("title")
	courseFilter.Description = g.Param("description")
	// Kurslarni filtrlash
	course, err := hand.Course.FilterCourses(courseFilter)
	if err != nil {
		g.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// Filtrlangan kurslarni JSON formatida qaytarish
	g.JSON(http.StatusOK, course)
}
