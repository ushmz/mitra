package database

import (
	"errors"
	"fmt"
	"mitra/config"
	"time"

	"github.com/jmoiron/sqlx"
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

	for i := 0; i < 30; i++ {
		if err := d.Ping(); err != nil {
			time.Sleep(1 * time.Second)
			continue
		}
		return d, nil
	}
	return nil, errors.New("Failed to connect DB: Connection timeout")
}
