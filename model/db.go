package model

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
)

// InitDB initialize and return DB connection
func InitDB() (*sqlx.DB, error) {
	hostname := os.Getenv("HOST")
	username := os.Getenv("USER")
	password := os.Getenv("PASS")
	database := os.Getenv("DATABASE")
	port := os.Getenv("PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		username,
		password,
		hostname,
		port,
		database,
	)

	d, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	for i := 0; i < 30; i++ {
		if err := d.Ping(); err != nil {
			time.Sleep(1 * time.Second)
			continue
		}
		return d, nil
	}
	return nil, errors.New("Failed to connect DB: Connection timeout")
}
