package storage

import (
	"database/sql"
	api "n11/Firdavs/dars2.17/handler"
	"n11/Firdavs/dars2.17/model"

	"github.com/gin-gonic/gin"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{db: db}
}

func  GetAll(c *gin.Context) {

	f := model.Filter{}

	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
	)

	query := "select id, first_name, last_name, age, gender, nation, field, parent_name, city from person "

	filter := "where true  "

	if len(f.Gender) > 0 {
		params["gender"] = f.Gender
		filter += " and gender = :gender "
	}

	if len(f.Nation) > 0 {
		params["nation"] = f.Gender
		filter += " and nation = :nation "
	}

	if len(f.Field) > 0 {
		params["field"] = f.Gender
		filter += " and field = :field "
	}

	if f.Age > 0 {
		params["age"] = f.Gender
		filter += " and age = :age "
	}

	if f.Limit > 0 {
		params["limit"] = f.Limit
		limit = ` LIMIT :limit`
	}

	//if f.Offset > 0 {
	//	params["offset"] = (f.Offset - 1) * f.Limit
	//	offset = ` OFFSET :offset`
	//}
	//
	query = query + filter + limit // + offset

	query, arr = api.ReplaceQueryParams(query, params)
	db := sql.DB
	db.Query(query, arr...)

}
