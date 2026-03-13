package storage

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type PostgresStore struct {
	DB *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	store := &PostgresStore{DB: db}

	if err := store.createTable(); err != nil {
		return nil, err
	}

	return store, nil
}

func (s *PostgresStore) createTable() error {

	query := `
	CREATE TABLE IF NOT EXISTS games (
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		developer TEXT NOT NULL,
		genre TEXT,
		release_year INT,
		difficulty INT,
		platform TEXT,
		boss_count INT
	);`

	_, err := s.DB.Exec(query)
	return err
}
