package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() error {
	db, err := sql.Open("mysql", "username:password@tcp(host:port)/database_name")
	if err != nil {
		return err
	}
	DB = db
	return nil
}
