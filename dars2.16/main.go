package main

import (
	handler "new/hendler"
	"new/postgres"
	"log"

	_ "github.com/lib/pq"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	userRepo := postgres.NewUserRepo(db)
	problemRepo := postgres.NewProblemRepo(db)
	solvedProblemRepo := postgres.NewSolvedProblemRepo(db)

	h := handler.Handler{
		User:          userRepo,
		Problem:       problemRepo,
		SolvedProblem: solvedProblemRepo,
	}

	router := gin.Default()

	// User routes
	router.GET("/user/:id", h.Userr)
	router.POST("/user/create", h.UserCreate)
	router.PUT("/user/update/:id", h.UserUpdate)
	router.DELETE("/user/delete/:id", h.UserDelete)

	// Problem routes
	router.GET("/problem/:id", h.Problem1)
	router.POST("/problem/create", h.ProblemCreate)
	router.PUT("/problem/update/:id", h.ProblemUpdate)
	router.DELETE("/problem/delete/:id", h.ProblemDelete)

	// SolvedProblem routes
	router.GET("/solvedproblem/:id", h.SolvedProblemGet)
	router.POST("/solvedproblem/create", h.SolvedProblemCreate)
	router.PUT("/solvedproblem/update/:id", h.SolvedProblemUpdate)
	router.DELETE("/solvedproblem/delete/:id", h.SolvedProblemDelete)

	log.Println("Server started at :8080")
	log.Fatal(router.Run(":8080"))
}
