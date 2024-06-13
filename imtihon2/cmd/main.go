package main

import (
	"n11/Firdavs/imtihon2/handler"
	"n11/Firdavs/imtihon2/storage/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	user := postgres.NewUser(db)
	courses := postgres.NewCourses(db)
	lesson := postgres.NewLesson(db)
	enrollments := postgres.NewEnrollments(db)

	l := handler.GIN(db, user, courses, lesson, enrollments)
	l.Run("localhost:8080")
}
