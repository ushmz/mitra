package database

import (
	"context"
	"errors"
	"fmt"
	"mitra/config"
	"time"

	"github.com/jmoiron/sqlx"
)

const (
	connectionAttempts = 3
	pingTimeoutSecs    = 10
)

// InitDB initialize and return DB connection
func InitDB() (*sqlx.DB, error) {
	conf := config.GetConfig()

	host := conf.GetString("mysql.host")
	user := conf.GetString("mysql.user")
	port := conf.GetString("mysql.port")
	pass := conf.GetString("mysql.password")
	database := conf.GetString("mysql.database")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		user,
		pass,
		host,
		port,
		database,
	)

	d, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	for i := 0; i < connectionAttempts; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), pingTimeoutSecs*time.Second)
		defer cancel()
		if err := d.PingContext(ctx); err != nil {
			if i == connectionAttempts-1 {
				// Return error or Panic (log.Fatal()) migh be better.
				return nil, errors.New("Failed to connect DB: Connection timeout")
			}
			fmt.Printf("Failed to ping DB: Retry in %d seconds", pingTimeoutSecs)
			time.Sleep(pingTimeoutSecs * time.Second)
		}
	}
	return d, nil
}
