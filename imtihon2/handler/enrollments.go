package handler

import (
	"fmt"
	"n11/Firdavs/imtihon2/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (hand *Handler) CreateEnrollments(g *gin.Context) {
  var enrollment model.Enrollments
  err := g.BindJSON(&enrollment)
  if err != nil {
    g.String(http.StatusBadRequest, err.Error())
  }
  err = hand.Enrollment.CreateEnrollments(&enrollment)
  if err != nil {
    g.String(http.StatusInternalServerError, err.Error())
  }
  g.JSON(http.StatusCreated, enrollment)
}


func (hand *Handler) UpdateEnrollment(g *gin.Context) {
  var enrollment model.Enrollments
  err := g.Bind(&enrollment)
  if err != nil {
    g.String(http.StatusBadRequest, err.Error())
  }
  err = hand.Enrollment.UpdateEnrollment(&enrollment)
  if err != nil {
    g.String(http.StatusInternalServerError, err.Error())
  }

  g.String(200, "SUCCESS")

}

func (hand *Handler) GetAllEnrollments(g *gin.Context) {
  courses, err := hand.Enrollment.ReadAllEnrollments()

  if err != nil {
    g.String(http.StatusInternalServerError, err.Error())
  }
  g.JSON(http.StatusOK, courses)

}

func (hand *Handler) GetEnrollmentById(g *gin.Context) {
  id := g.Param("enrollment_id")
  if id == "" {
    g.String(http.StatusBadRequest, "There is an error entering ID")
  }

  enrollment, err := hand.Enrollment.GetEnrollment(id)
  if err != nil {
	fmt.Println("64 hand", err)
    g.String(http.StatusInternalServerError, err.Error())
  }
  g.JSON(http.StatusOK, enrollment)

}

func (hand *Handler) DeleteEnrollment(g *gin.Context) {
  id := g.Param("enrollment_id")
  if id == "" {
    g.String(http.StatusBadRequest, "There is an error entering ID")
  }

  err := hand.Enrollment.DeleteEnrollment(id)
  if err != nil {
    g.String(http.StatusInternalServerError, err.Error())
  }
  g.String(200, "SUCCESS")

}

func (hand *Handler) GetFilterEnrollments(g *gin.Context) {
  EnrollmentFilter := model.Enrollments{}
  EnrollmentFilter.EnrollmentId = g.Param("enrollment_id")
  EnrollmentFilter.UserId = g.Param("user_id")
  EnrollmentFilter.CourseId = g.Param("course_id")
  EnrollmentFilter.EnrollmentDate = g.Param("enrollment_date")

  g.JSON(http.StatusOK, EnrollmentFilter)
}
