package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

type newSheet struct {
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
	Templates string `json:"templates"`
}

func (app *application) addSheet(w http.ResponseWriter, r *http.Request) {
	var s newSheet

	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = app.sheets.Insert(s.Name, s.ShortName, s.Templates)
	if err != nil {
		app.serveError(w, r, err)
		return
	}

	resp := make(map[string]string)
	resp["message"] = fmt.Sprintf("Sheet named: %s added successfully", s.Name)

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		return
	}
}

func (app *application) updateSheet(w http.ResponseWriter, r *http.Request) {
	var ns newSheet

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&ns)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = app.sheets.Update(id, ns.Name, ns.ShortName, ns.Templates)
	if err != nil {
		app.serveError(w, r, err)
		return
	}

	resp := make(map[string]string)
	resp["message"] = "sheet edited successfully"

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		return
	}

}

func (app *application) getSheetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	s, err := app.sheets.GetByID(id)
	if err != nil {
		app.serveError(w, r, err)
		return
	}

	err = json.NewEncoder(w).Encode(s)
	if err != nil {
		return
	}

}

type newUser struct {
	UserID    string `json:"firebase_uuid"`
	Name      string `json:"name"`
	ManNumber int    `json:"man_number"`
	Picture   string `json:"picture"`
	Email     string `json:"email"`
}

func (app *application) addUser(w http.ResponseWriter, r *http.Request) {
	var nu newUser

	err := json.NewDecoder(r.Body).Decode(&nu)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = app.users.Insert(nu.ManNumber, nu.Name, nu.Picture, nu.Email, nu.UserID)
	if err != nil {
		var e *pgconn.PgError
		if errors.As(err, &e) && e.Code == pgerrcode.UniqueViolation {
			http.Error(w, "User has already been added.", http.StatusBadRequest)
			return
		}
		app.serveError(w, r, err)
		return
	}

	resp := make(map[string]string)
	resp["message"] = fmt.Sprintf("User named: %s added successfully", nu.Name)

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		return
	}
}

type newUserRole struct {
	UserID   int    `json:"user_id"`
	RoleName string `json:"role_name"`
}

func (app *application) addUserRole(w http.ResponseWriter, r *http.Request) {
	var role newUserRole

	sheetID, err := strconv.Atoi(r.PathValue("sheet_id"))
	if err != nil || sheetID < 1 {
		http.NotFound(w, r)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&role)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = app.roles.Insert(role.UserID, sheetID, role.RoleName)
	if err != nil {
		app.serveError(w, r, err)
		return
	}

	resp := make(map[string]string)
	resp["message"] = "User role added successfully"

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		return
	}
}

func (app *application) removeUserRole(w http.ResponseWriter, r *http.Request) {
	var role newUserRole

	sheetID, err := strconv.Atoi(r.PathValue("sheet_id"))
	if err != nil || sheetID < 1 {
		http.NotFound(w, r)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&role)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := app.roles.Delete(role.UserID, sheetID)
	if err != nil {
		app.serveError(w, r, err)
		return
	}

	fmt.Println(result.RowsAffected())

	resp := make(map[string]string)
	resp["message"] = "User role removed successfully"

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		return
	}

}

// func (app *application) editUserRole(w http.ResponseWriter, r *http.Request) {

//   if app.auth.CheckAuthLevel()
// }
