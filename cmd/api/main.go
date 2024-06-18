package main

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"platinum_grid/internal/models"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
)

var (
	database = os.Getenv("DB_DATABASE")
	password = os.Getenv("DB_PASSWORD")
	username = os.Getenv("DB_USERNAME")
	db_port  = os.Getenv("DB_PORT")
	host     = os.Getenv("DB_HOST")
	addr     = os.Getenv("PORT")
)

type application struct {
	logger *slog.Logger
	sheets *models.SheetModel
	users  *models.UserModel
}

func main() {

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, db_port, database)

	fmt.Println(connStr)
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}))

	db, err := openDB(connStr)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close()

	app := &application{
		logger: logger,
		sheets: &models.SheetModel{DB: db},
		users:  &models.UserModel{DB: db},
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", addr),
		Handler: app.routes(),
	}

	logger.Info("starting server", "addr", addr)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
