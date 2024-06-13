package handler

import (
	"fmt"
	"n11/Firdavs/imtihon2/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (hand *Handler) CreateCurse(g *gin.Context) {
  var course model.Courses
  err := g.BindJSON(&course)
  if err != nil {
    g.String(http.StatusBadRequest, err.Error())
  }
  err = hand.Course.CreateCurse(&course)
  if err != nil {
    g.String(http.StatusInternalServerError, err.Error())
  }
  g.JSON(http.StatusCreated, course)
}


func (hand *Handler) UpdateCourse(g *gin.Context) {
  var course model.Courses
  err := g.Bind(&course)
  if err != nil {
    g.String(http.StatusBadRequest, err.Error())
  }
  err = hand.Course.UpdateCourse(&course)
  if err != nil {
    g.String(http.StatusInternalServerError, err.Error())
  }

  g.String(200, "SUCCESS")

}

func (hand *Handler) GetAllCourses(g *gin.Context) {
  courses, err := hand.Course.ReadAllCourses()

  if err != nil {
    g.String(http.StatusInternalServerError, err.Error())
  }
  g.JSON(http.StatusOK, courses)

}

func (hand *Handler) GetCourseById(g *gin.Context) {
  id := g.Param("course_id")
  if id == "" {
    g.String(http.StatusBadRequest, "There is an error entering ID")
  }

  course, err := hand.Course.ReadCourse(id)
  if err != nil {
	fmt.Println("64 hand", err)
    g.String(http.StatusInternalServerError, err.Error())
  }
  g.JSON(http.StatusOK, course)

}

func (hand *Handler) DeleteCourse(g *gin.Context) {
  id := g.Param("course_id")
  if id == "" {
    g.String(http.StatusBadRequest, "There is an error entering ID")
  }

  err := hand.Course.DeleteCourse(id)
  if err != nil {
    g.String(http.StatusInternalServerError, err.Error())
  }
  g.String(200, "SUCCESS")

}

func (hand *Handler) GetFilterCourses(g *gin.Context) {
  courseFilter := model.Courses{}
  courseFilter.Course_id = g.Param("course_id")
  courseFilter.Title = g.Param("title")
  courseFilter.Description = g.Param("description")

  g.JSON(http.StatusOK, courseFilter)
}
