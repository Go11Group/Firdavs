package postgres

import (
	"n11/Firdavs/dars2.9/model"
	"strconv"

	"gorm.io/gorm"
)

type UserRepo struct {
	db gorm.DB
}

func (u *UserRepo) CreateTable(user *model.Users) error {
	err := u.db.AutoMigrate(&model.Users{})

	if err != nil {
		return  err
	}
	return nil
}

func (u *UserRepo) Create(user *model.Users)  {

	for i := 1; i <= 25; i++ {
		user := model.Users{
			FirstName:  "Ali" + strconv.Itoa(i),
			LastName:   "Boqiyev",
			Email:      "wf" + strconv.Itoa(i) + "@kwej.das",
			Password:   "12" + strconv.Itoa(i) + "23",
			Age:        25,
			Field:      "Golang developer",
			Gender:     "Male",
			IsEmployee: true,
		}
		u.db.Create(&user)

	}
}


