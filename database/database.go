package database

import (
	"context"
	"errors"
	"fmt"
	"log"
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
	log.Println(username)
	dbURL := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s", username, password, host, port, dbname, sslmode)
	// dbURL := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s", "postgres", "password", "localhost", "5432", "sle", "disable") // only for deb mode

	dbPoolConfig, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		return nil, errors.New("error during opening the database")
	}

	db, err := pgxpool.NewWithConfig(context.Background(), dbPoolConfig)
	if err != nil {
		log.Fatalln("Unable to create connection pool:", err)
	}

	if err := db.Ping(context.Background()); err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

func (d *Database) GetDB() *pgxpool.Pool {
	return d.db
}

func (d *Database) CloseDB() {
	d.db.Close()
}
