package models

import (
	"database/sql"
)

type Sheet struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
	Templates string `json:"templates"`
}

type SheetModel struct {
	DB *sql.DB
}

func (m *SheetModel) Insert(name, shortName, templates string) error {

	stmt := `INSERT INTO sheets (name, short_name, templates) VALUES ($1, $2, $3) RETURNING *;`

	_, err := m.DB.Exec(stmt, name, shortName, templates)
	if err != nil {
		return err
	}

	return nil
}

func (m *SheetModel) Update(id int, name, shortName, templates string) error {

	stmt := `UPDATE SHEETS SET
  name = COALESCE(NULLIF($2,''), name),
  short_name = COALESCE(NULLIF($3,''), short_name),
  templates = COALESCE(NULLIF($4,''), templates)
  WHERE id = $1
  RETURNING *;
  `

	_, err := m.DB.Exec(stmt, id, name, shortName, templates)
	if err != nil {
		return err
	}

	return nil
}

func (m *SheetModel) GetByID(id int) (Sheet, error) {
	var s Sheet

	stmt := `SELECT * FROM sheets WHERE id = $1 LIMIT 1;`

	err := m.DB.QueryRow(stmt, id).Scan(&s.ID, &s.Name, &s.ShortName, &s.Templates)
	if err != nil {
		return Sheet{}, err
	}

	return s, nil
}
