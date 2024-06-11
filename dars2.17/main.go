package main

import (
	storage "n11/Firdavs/dars2.17/storage/postgres"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/getAll", storage.GetAll)

	r.Run(":8080")
}
