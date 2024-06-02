package postgres

import (
	"fmt"
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
		return err
	}
	return nil
}

func (u *UserRepo) Create(user *model.Users) error {

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
		err := u.db.Create(&user).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (u *UserRepo) Update(user model.Users) error {
	updates := map[string]interface{}{
		"first_name":  user.FirstName,
		"last_name":   user.LastName,
		"Email":       user.Email,
		"Password":    user.Password,
		"age":         user.Age,
		"field":       user.Field,
		"is_employee": user.IsEmployee,
		"gender":      user.Gender,
	}

	if err := u.Db.Model(&user).Where("id = ?", user.ID).Updates(updates).Error; err != nil {
		return fmt.Errorf("failed while updating user %w", err)
	}

	return nil
}
