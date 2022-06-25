package db

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	database_name := os.Getenv("DB_NAME")
	dbconf := user + ":" + pass + "@tcp(" + host + ")/" + database_name + "?charset=utf8mb4&parseTime=true&loc=Asia%2FTokyo"

	if db, err := sql.Open("mysql", dbconf); err != nil {
		panic(err)
	} else {
		// See "Important settings" section.
		db.SetConnMaxLifetime(time.Minute * 3)
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(10)

		return db
	}
}
