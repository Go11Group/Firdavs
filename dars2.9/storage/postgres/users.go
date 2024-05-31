package postgres

import (
	"n11/Firdavs/dars2.9/model"

	"gorm.io/gorm"
)

type UserRepo struct {
	db gorm.DB
}

func (u *UserRepo) CreateTable(user *model.Users) error {
	u.db.AutoMigrate(&model.Users{})

	return nil
}

func (u *UserRepo) Create(user model.Users)  {

	u.db.Create(&model.Users{FirstName: })
}
