package models

import (
	"database/sql"
)

type UserRoles struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id"`
	RoleName string `json:"role_name"`
	SheetID  int    `json:"sheet_id"`
}

type UserRoleModel struct {
	DB *sql.DB
}

func (m *UserRoleModel) Insert(userID int, sheetID int, roleName string) error {
	stmt := `INSERT INTO user_roles (user_id, role_name, sheet_id) VALUES ($1, $2, $3) RETURNING *;`

	_, err := m.DB.Exec(stmt, userID, roleName, sheetID)
	if err != nil {
		return err
	}

	return nil
}

func (m *UserRoleModel) Delete(userId, sheetId int) (sql.Result, error) {
	stmt := `DELETE FROM user_roles WHERE user_id = $1 AND sheet_id = $2;`

	result, err := m.DB.Exec(stmt, userId, sheetId)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type UsersRoles struct {
	Users []UserRoleModel
}

func (m *UserRoleModel) Update(users UsersRoles, sheet_id, user_id int, roleName string) (sql.Result, error) {

	stmt := `UPDATE user_roles SET role_name = $1 WHERE user_id = $2 AND sheet_id = $3`

	result, err := m.DB.Exec(stmt, roleName, user_id, sheet_id)
	if err != nil {
		return nil, err
	}

	return result, nil
}
