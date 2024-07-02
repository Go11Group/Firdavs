package handler

import (
	"new/postgres"
)

type Handler struct {
	User          *postgres.UserRepo
	Problem       *postgres.ProblemRepo
	SolvedProblem *postgres.SolvedProblemRepo
}
