package main

import (
	"n11/Firdavs/imtihon2/handler"
	"n11/Firdavs/imtihon2/storage/postgres"
)

func main() {
	// Bazaga ulanish
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Repozitoriyalarni boshlash
	userRepo := postgres.NewUser(db)
	coursesRepo := postgres.NewCourses(db)
	lessonRepo := postgres.NewLesson(db)
	enrollmentsRepo := postgres.NewEnrollments(db)
	fivetaskRepo := postgres.NewFivetask(db)

	// GIN xandlerni ishlatib serverni boshlash
	server := handler.GIN(db, userRepo, coursesRepo, lessonRepo, enrollmentsRepo, fivetaskRepo)
	server.Run("localhost:8080")
}
