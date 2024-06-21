package models

import (
	"database/sql"
	"strings"
)

type AuthModel struct {
	DB *sql.DB
}

type AuthObj struct {
	Read        bool
	Write       bool
	UserManage  bool
	SheetManage bool
}

var authObj = map[string]AuthObj{
	"owner":  {Read: true, Write: true, UserManage: true, SheetManage: true},
	"viewer": {Read: true, Write: false, UserManage: false, SheetManage: false},
	"editor": {Read: true, Write: true, UserManage: false, SheetManage: false},
}

func (m *AuthModel) requestCurrentUser(id string) User {
	var u User
	stmt := `SELECT * FROM users WHERE firebase_uuid = $1;`

	result := m.DB.QueryRow(stmt, id)
	err := result.Scan(&u)
	if err != nil {
		return User{}
	}
	return u
}

type CheckAuth struct {
	UserId   int    `json:"user_id"`
	RoleName string `json:"role_name"`
	SheetId  int    `json:"sheet_id"`
}

func (m *AuthModel) CheckAuthLevel(action, id string, sheet int) (bool, error) {
	var a CheckAuth

	currentUser := m.requestCurrentUser(id)

	stmt := `SELECT * FROM user_roles WHERE user_id = $1 AND sheet_id = $2`

	err := m.DB.QueryRow(stmt, currentUser.ID, sheet).Scan(&a)
	if err != nil {
		return false, err
	}

	userRole := strings.ToLower(a.RoleName)

	switch action {
	case "read":
		return authObj[userRole].Read, nil
	case "write":
		return authObj[userRole].Write, nil
	case "userManage":
		return authObj[userRole].UserManage, nil
	case "sheetManage":
		return authObj[userRole].SheetManage, nil
	default:
		return false, nil
	}
}
