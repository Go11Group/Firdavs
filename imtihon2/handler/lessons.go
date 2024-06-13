package handler

import (
	"fmt"
	"n11/Firdavs/imtihon2/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (hand *Handler) CreateLesson(g *gin.Context) {
	var lesson model.Lessons
	err := g.BindJSON(&lesson)
	if err != nil {
		g.String(http.StatusBadRequest, err.Error())
	}
	err = hand.Lesson.CreateLesson(&lesson)
	if err != nil {
		g.String(http.StatusInternalServerError, err.Error())
	}
	g.JSON(http.StatusCreated, lesson)
}



func (hand *Handler) UpdateLesson(g *gin.Context) {
	var lesson model.Lessons
	err := g.Bind(&lesson)
	if err != nil {
		g.String(http.StatusBadRequest, err.Error())
	}
	err = hand.Lesson.UpdateLesson(&lesson)
	if err != nil {
		g.String(http.StatusInternalServerError, err.Error())
	}

	g.String(200, "SUCCESS")

}

func (hand *Handler) GetAllLessons(g *gin.Context) {
	courses, err := hand.Lesson.GetAllLessons()

	if err != nil {
		g.String(http.StatusInternalServerError, err.Error())
	}
	g.JSON(http.StatusOK, courses)

}

func (hand *Handler) GetLessonById(g *gin.Context) {
	id := g.Param("lesson_id")
	if id == "" {
		g.String(http.StatusBadRequest, "There is an error entering ID")
	}

	lesson, err := hand.Lesson.ReadLesson(id)
	if err != nil {
		fmt.Println("64 hand", err)
		g.String(http.StatusInternalServerError, err.Error())
	}
	g.JSON(http.StatusOK, lesson)

}

func (hand *Handler) DeleteLesson(g *gin.Context) {
	id := g.Param("lesson_id")
	if id == "" {
		g.String(http.StatusBadRequest, "There is an error entering ID")
	}

	err := hand.Lesson.DeleteLesson(id)
	if err != nil {
		g.String(http.StatusInternalServerError, err.Error())
	}
	g.String(200, "SUCCESS")

}

func (hand *Handler) GetFilterLessons(g *gin.Context) {
	LessonFilter := model.Lessons{}
	LessonFilter.LessonId = g.Param("lesson_id")
	LessonFilter.CourseId = g.Param("course_id")
	LessonFilter.Title = g.Param("title")
	LessonFilter.Content = g.Param("content")

	g.JSON(http.StatusOK, LessonFilter)
}
