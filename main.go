package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/tylerbrewer2/wedding-backend/config"
	"github.com/tylerbrewer2/wedding-backend/rsvps"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := startAndVerifyDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Starting web server")

	rsvps.RegisterRoutes(db)

	log.Fatal(http.ListenAndServe(":8080", nil))
	fmt.Println("Server shutting down")
}

func startAndVerifyDB(cfg config.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s dbname=%s sslmode=disable", cfg.DB.Username, cfg.DB.Name)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, errors.Wrap(err, "error opening conn to database")
	}

	err = db.Ping()
	return db, errors.Wrap(err, "failed to ping database at startup")
}
