package handler

import (
	"fmt"
	"github.com/Go11Group/at_lesson/lesson43/metro_service/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *handler) CreateStation(ctx *gin.Context) {
	stn := models.CreateStation{}

	err := ctx.ShouldBindJSON(&stn)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	err = h.Station.Create(&stn)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, "OKAY")
}
