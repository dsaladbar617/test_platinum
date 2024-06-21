package models

import "database/sql"

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
