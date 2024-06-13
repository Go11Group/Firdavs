package handler

import (
	"fmt"
	"n11/Firdavs/imtihon2/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (hand *Handler) CreateUser(g *gin.Context) {
  var user model.User
  err := g.BindJSON(&user)
  if err != nil {
    g.String(http.StatusBadRequest, err.Error())
  }
  err = hand.User.CreateUser(&user)
  if err != nil {
    g.String(http.StatusInternalServerError, err.Error())
  }
  g.JSON(http.StatusCreated, user)
}
// TODO
// func (hand *Handler) GetCourseByUsers(g *gin.Context) {
//   userId, course, err := hand.User.GetCourseByUser(g.Param("id"))
//   if err != nil {
//     g.String(http.StatusInternalServerError, err.Error())
//   }
//   g.JSON(http.StatusOK,userId, course)

// }

func (hand *Handler) UpdateUser(g *gin.Context) {
  var user model.User
  err := g.Bind(&user)
  if err != nil {
    g.String(http.StatusBadRequest, err.Error())
  }
  err = hand.User.UpdateUser(&user)
  if err != nil {
    g.String(http.StatusInternalServerError, err.Error())
  }

  g.String(200, "SUCCESS")

}

func (hand *Handler) GetAllUsers(g *gin.Context) {
  users, err := hand.User.ReadAllUsers()
  if err != nil {
    g.String(http.StatusInternalServerError, err.Error())
  }
  g.JSON(http.StatusOK, users)

}

func (hand *Handler) GetUserById(g *gin.Context) {
  id := g.Param("user_id")
  if id == "" {
    g.String(http.StatusBadRequest, "There is an error entering ID")
  }

  user, err := hand.User.ReadUser(id)
  if err != nil {
	fmt.Println("he;;d", err)
    g.String(http.StatusInternalServerError, err.Error())
  }
  g.JSON(http.StatusOK, user)

}

func (hand *Handler) DeleteUsers(g *gin.Context) {
  id := g.Param("user_id")
  if id == "" {
    g.String(http.StatusBadRequest, "There is an error entering ID")
  }

  err := hand.User.DeleteUser(id)
  if err != nil {
    g.String(http.StatusInternalServerError, err.Error())
  }
  g.String(200, "SUCCESS")

}

func (hand *Handler) GetFilterUsers(g *gin.Context) {
  userFilter := model.User{}
  userFilter.Id = g.Param("user_id")
  userFilter.Name = g.Param("name")
  userFilter.Email = g.Param("email")
  userFilter.Birthday = g.Param("birthday")
  userFilter.Password = g.Param("password")

  g.JSON(http.StatusOK, userFilter)
}
