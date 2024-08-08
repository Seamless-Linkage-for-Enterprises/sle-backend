package database

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	db *pgxpool.Pool
}

func NewDatabase() (*Database, error) {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	dbURL := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s", username, password, host, port, dbname, sslmode)
	dbpool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		return nil, errors.New("error during opening the database")
	}
	defer dbpool.Close()

	if err := dbpool.Ping(context.Background()); err != nil {
		return nil, err
	}

	return &Database{db: dbpool}, nil
}
