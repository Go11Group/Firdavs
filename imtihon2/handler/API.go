package handler

import (
	"database/sql"
	"n11/Firdavs/imtihon2/storage/postgres"

	"github.com/gin-gonic/gin"
)

func GIN(db *sql.DB, user *postgres.User, course *postgres.Courses) *gin.Engine {
	r := gin.Default()

	handler := NewHandler(db, user, course)

	r.GET("/user", handler.GetAllUsers)
	r.GET("/user/:user_id", handler.GetUserById)
	r.POST("/user", handler.CreateUser)
	r.PUT("/user/:id", handler.UpdateUser)
	r.DELETE("/user/:id", handler.DeleteUsers)

	return r
}
