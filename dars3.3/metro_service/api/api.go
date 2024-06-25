package api

import (
	"database/sql"
	"github.com/Go11Group/at_lesson/lesson43/metro_service/api/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routes(db *sql.DB) *http.Server {
	mux := gin.Default()

	h := handler.NewHandler(db)

	mux.POST("station/create", h.CreateStation)

	return &http.Server{Handler: mux, Addr: ":8080"}
}
