package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"n11/Firdavs/imtihon2/model"
	"time"

	"github.com/google/uuid"
)

// User strukturasi, DB ulanishini o'z ichiga olgan foydalanuvchi repozitoriyasi.
type User struct {
	db *sql.DB
}

// NewUser funksiya, berilgan DB ulanishi bilan yangi foydalanuvchi repozitoriyasini yaratadi.
func NewUser(db *sql.DB) *User {
	return &User{db}
}

// CreateUser funksiya, yangi foydalanuvchi yozuvinini bazaga qo'shadigan funksiya.
func (u *User) CreateUser(user *model.User) error {
	birthday, err := time.Parse("2006-01-02", user.Birthday)
	if err != nil {
		fmt.Println("Tug'ilgan kunni tahlil qilishda xatolik:", err)
		return errors.New("tug'ilgan kunni tahlil qilishda xatolik")
	}

	user.Id = uuid.NewString()
	_, err = u.db.Exec("insert into users (user_id, name, email, birthday, password) VALUES ($1, $2, $3, $4, $5)",
		user.Id, user.Name, user.Email, birthday, user.Password)
	if err != nil {
		fmt.Println("Foydalanuvchini qo'shishda xatolik:", err)
		return err
	}
	return nil
}

// ReadUser funksiya, foydalanuvchi ID bo'yicha foydalanuvchi yozuvinini o'qish uchun ishlatiladi.
func (u *User) ReadUser(id string) (*model.User, error) {
	row := u.db.QueryRow("select * from users where user_id = $1", id)

	var user model.User
	err := row.Scan(
		&user.Id, &user.Name, &user.Email, &user.Birthday, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeleteAt)
	if err != nil {
		fmt.Println("Foydalanuvchini o'qishda xatolik:", err)
		return nil, err
	}
	return &user, nil
}

// UpdateUser funksiya, mavjud foydalanuvchi yozuvinini bazada yangilash uchun ishlatiladi.
func (u *User) UpdateUser(user *model.User) error {
	birthday, err := time.Parse("2006-01-02", user.Birthday)
	if err != nil {
		fmt.Println("Tug'ilgan kunni tahlil qilishda xatolik:", err)
		return errors.New("tug'ilgan kunni tahlil qilishda xatolik")
	}
	_, err = u.db.Exec("update users set name = $1, email = $2, birthday = $3, password = $4 where user_id = $5",
		user.Name, user.Email, birthday, user.Password, user.Id)
	if err != nil {
		fmt.Println("Foydalanuvchini yangilashda xatolik:", err)
		return err
	}
	return nil
}

// DeleteUser funksiya, foydalanuvchi ID bo'yicha foydalanuvchi yozuvinini bazadan o'chirish uchun ishlatiladi.
func (u *User) DeleteUser(id string) error {
	_, err := u.db.Exec("delete from users where user_id = $1", id)
	if err != nil {
		fmt.Println("Foydalanuvchini o'chirishda xatolik:", err)
		return err
	}
	return nil
}

// ReadAllUsers funksiya, barcha foydalanuvchi yozuvinini bazadan o'qish uchun ishlatiladi.
func (u *User) ReadAllUsers() ([]*model.User, error) {
	rows, err := u.db.Query("select * from users")
	if err != nil {
		fmt.Println("Barcha foydalanuvchilarni o'qishda xatolik:", err)
		return nil, err
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(
			&user.Id, &user.Name, &user.Email, &user.Birthday, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeleteAt)
		if err != nil {
			fmt.Println("Foydalanuvchi qatordan o'qishda xatolik:", err)
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

// FilterUsers funksiya, filtr parametrlariga asosan foydalanuvchi yozuvinini bazadan o'qish uchun ishlatiladi.
func (u *User) FilterUsers(f model.FilterUsers) ([]model.User, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
	)
	query := `select user_id, name, email, birthday, password
  from users `
	filter := ` where true`

	if len(f.Id) > 0 {
		params["user_id"] = f.Id
		filter += " and user_id = :user_id "
	}
	if len(f.Name) > 0 {
		params["name"] = f.Name
		filter += " and name = :name "
	}

	if len(f.Email) > 0 {
		params["email"] = f.Email
		filter += " and email = :email "
	}

	if len(f.Birthday) > 0 {
		params["birthday"] = f.Birthday
		filter += " and birthday = :birthday "
	}

	if len(f.Password) > 0 {
		params["password"] = f.Password
		filter += " and password = :password "
	}

	query = query + filter

	query, arr = ReplaceQueryParams(query, params)
	rows, err := u.db.Query(query, arr...)
	if err != nil {
		fmt.Println("Foydalanuvchilarni filtrlashda xatolik:", err)
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Birthday, &user.Password)
		if err != nil {
			fmt.Println("Foydalanuvchi qatordan o'qishda xatolik:", err)
			return nil, err
		}

		users = append(users, user)
	}
	return users, nil
}
