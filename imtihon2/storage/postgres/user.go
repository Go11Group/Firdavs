package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"n11/Firdavs/imtihon2/model"
	"time"

	"github.com/google/uuid"
)

type User struct {
	db *sql.DB
}

func NewUser(db *sql.DB) *User {
	return &User{db}
}

func (u *User) CreateUser(user *model.User) error {
	birthday, err := time.Parse("2006-01-02", user.Birthday)
	if err != nil {
		fmt.Println("2hnn2", err)
		return errors.New("failed to parse birthday")
	}

	user.Id = uuid.NewString()
	_, err = u.db.Exec("insert into users (user_id,name,email,birthday,password) VALUES ($1,$2,$3,$4,$5)",
		user.Id, user.Name, user.Email, birthday, user.Password)
	if err != nil {

		fmt.Println("3hmmmmnn2", err)
		return err
	}
	return nil
}

func (u *User) GetCourseByUser(userId string) (string, []model.Courses, error) {
	var courses string
	err := u.db.QueryRow("select course_id from enrollments where user_id = $1", userId).Scan(&courses)
	if err != nil {
		return "", nil, err
	}
	rows, err := u.db.Query("select id,title,description from  courses where id=$1", &courses)
	if err != nil {
		return "", nil, err
	}
	course := []model.Courses{}
	for rows.Next() {
		cour := model.Courses{}
		err := rows.Scan(&cour.Course_id, &cour.Title, &cour.Description)
		if err != nil {
			return "", nil, err
		}
		course = append(course, cour)
	}
	return userId, course, nil
}

func (u *User) ReadUser(id string) (*model.User, error) {
	row := u.db.QueryRow("select * from users where user_id = $1", id)

	fmt.Println(id)
	var user model.User
	err := row.Scan(
		&user.Id, &user.Name, &user.Email, &user.Birthday, &user.Password, &user.CreatedAt, &user.UpdateAt, &user.DeleteAt)
	if err != nil {
		fmt.Println("64", err)
		return nil, err
	}
	return &user, nil
}

func (u *User) UpdateUser(user *model.User) error {
	birthday, err := time.Parse("2006-01-02", user.Birthday)
	if err != nil {
		return errors.New("failed to parse birthday")
	}
	_, err = u.db.Exec("update users set name = $1, email = $2, birthday = $3, password = $4 where user_id = $5",
		user.Name, user.Email, birthday, user.Password, user.Id)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) DeleteUser(id string) error {
	_, err := u.db.Exec("delete from users where user_id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) ReadAllUsers() ([]*model.User, error) {
	rows, err := u.db.Query("select * from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []*model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(
			&user.Id, &user.Name, &user.Email, &user.Birthday, &user.Password, &user.CreatedAt, &user.UpdateAt, &user.DeleteAt)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (u *User) FilterUsers(f model.FilterUsers) ([]model.User, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
	)
	query := `select user_id,name,email,birthday,password
  from users `
	filter := ` where true`

	if len(f.Id) > 0 {
		params["user_id"] = f.Id
		filter += " and user_id = :user_id "
	}
	if len(f.Name) > 0 {
		params["name"] = f.Name
		filter += "and name = :name "
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
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Birthday, &user.Password)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	return users, nil
}
