package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	http.HandleFunc("/save", saveData)
	http.HandleFunc("/read", readData)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func initDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS data (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            value TEXT
        );
    `)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func saveData(w http.ResponseWriter, r *http.Request) {
	value := r.URL.Query().Get("value")
	db, err := initDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO data (value) VALUES (?)", value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Data saved successfully")
}

func readData(w http.ResponseWriter, r *http.Request) {
	db, err := initDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT value FROM data")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var values []string
	for rows.Next() {
		var value string
		err := rows.Scan(&value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		values = append(values, value)
	}

	fmt.Fprintf(w, "Data: %v", values)
}
