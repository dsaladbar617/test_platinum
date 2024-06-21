package main

import (
	"net/http"

	"github.com/justinas/alice"
	"github.com/rs/cors"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	standard := alice.New(cors.Default().Handler)

	// // SHEET ROUTES
	mux.HandleFunc("POST /sheet", app.addSheet)
	mux.HandleFunc("PUT /edit_sheet/{id}", app.updateSheet)
	mux.HandleFunc("GET /get_sheet/{id}", app.getSheetByID)

	// USER ROUTES
	mux.HandleFunc("POST /add_user", app.addUser)
	mux.HandleFunc("POST /add_user_roles/{sheet_id}", app.addUserRole)

	// mux.HandleFunc("/health", s.healthHandler)

	root := http.NewServeMux()
	root.Handle("/api/", http.StripPrefix("/api", mux))

	return standard.Then(root)
}
