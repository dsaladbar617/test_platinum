package models

import (
	"database/sql"
)

type User struct {
	ID        int    `json:"id"`
	UserID    string `json:"firebase_uuid"`
	Name      string `json:"name"`
	ManNumber int    `json:"man_number"`
	Picture   string `json:"picture"`
	Email     string `json:"email"`
}

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(manNumber int, name, picture, email, userId string) error {
	stmt := `INSERT INTO users (firebase_uuid, name, picture, email) VALUES ($1, $2, $3, $4) RETURNING *;`

	_, err := m.DB.Exec(stmt, userId, name, picture, email)
	if err != nil {
		return err
	}

	return nil
}

func (m *UserModel) GetByID(id int) (User, error) {
	var u User

	stmt := `SELECT * FROM users WHERE id = $1 LIMIT 1;`

	err := m.DB.QueryRow(stmt, id).Scan(&u)
	if err != nil {
		return User{}, err
	}
	return u, nil
}
