package postgres

import (
	"database/sql"
	"n11/Firdavs/dars2.14/model"
	
)

type UserRepo struct {
	db *sql.DB
}

func User_NewRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) User_Get(id string) (model.User, error) {
	var user model.User
	err := u.db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(
		&user.Id, &user.Username, &user.Email, &user.Password)

	return user, err
}

func (u *UserRepo) User_delete(id string) error {
	_, err := u.db.Exec(`delete from users where id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepo) User_Create(user model.User) error {
	_, err := u.db.Exec(`INSERT INTO users (username, email, password) 
                       VALUES ($1, $2, $3)`,
		user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}
func (repo *UserRepo) User_Update(id int, user model.User) error {
	query := `UPDATE users SET username=$1, email=$2 WHERE id=$3`
	_, err := repo.db.Exec(query, user.Username, user.Email, id)
	if err != nil {
		return err
	}
	return nil
}
