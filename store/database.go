package store

import (
	"context"
	"errors"
	"fmt"
	"mitra/config"
	"time"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	connectionAttempts = 3
	pingTimeoutSecs    = 10
)

var (
	dialect = goqu.Dialect("mysql")
)

// InitDB initialize and return DB connection
func InitDB() (*sqlx.DB, error) {
	conf := config.GetConfig()

	config := conf.GetString("config")
	host := conf.GetString("mysql.host")
	user := conf.GetString("mysql.user")
	port := conf.GetString("mysql.port")
	pass := conf.GetString("mysql.password")
	database := conf.GetString("mysql.database")

	var dsn string

	if config == "local" {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
			user,
			pass,
			host,
			port,
			database,
		)
	} else {
		dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?tls=true&charset=utf8mb4&parseTime=True",
			user,
			pass,
			host,
			database,
		)
	}

	d, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	for i := 0; i < connectionAttempts; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), pingTimeoutSecs*time.Second)
		defer cancel()
		if err := d.PingContext(ctx); err != nil {
			if i == connectionAttempts-1 {
				// Return error or Panic (log.Fatal()) might be better.
				return nil, errors.New("Failed to connect DB: Connection timeout")
			}
			fmt.Printf("Failed to ping DB: Retry in %d seconds\n", pingTimeoutSecs)
			time.Sleep(pingTimeoutSecs * time.Second)
		}
	}
	return d, nil
}
