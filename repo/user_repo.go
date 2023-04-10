package repo

import (
	"database/sql"
	"errors"
	"foods/model"
)

type UserRepo interface {
	CreateUser(user *model.User) error
	GetUserByUsername(username string) (model.User, error)
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepo {
	return &userRepo{db}
}

func (u *userRepo) CreateUser(user *model.User) error {
	var id int
	query := `insert into users (username, password) values ($1, $2) returning id`
	err := u.db.QueryRow(query, user.Username, user.Password).Scan(&id)

	if err != nil {
		return err
	}

	user.ID = id
	return nil
}

func (u *userRepo) GetUserByUsername(username string) (model.User, error) {
	var user model.User
	query := `select id, username, password from users where username = $1`
	rows, err := u.db.Query(query, username)

	if err != nil {
		return model.User{}, err
	}

	for rows.Next() {
		rows.Scan(
			&user.ID,
			&user.Username,
			&user.Password,
		)
	}

	if user.Username == "" {
		return model.User{}, errors.New("")
	}

	return user, nil
}
