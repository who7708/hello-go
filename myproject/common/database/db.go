package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// InitDB mysql -uroot -proot1234 --default-character-set=utf8 -Dtest
func InitDB() error {
	db, err := sql.Open("mysql", "root:root1234@tcp(127.0.0.1:3305)/test")
	if err != nil {
		return err
	}
	DB = db
	return nil
}
