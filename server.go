package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func startServer(db *sql.DB) {
	http.HandleFunc("/unit", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var unit Unit
		err := json.NewDecoder(r.Body).Decode(&unit)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		insertUnit(db, unit)

		fmt.Fprintf(w, "Success insertUnit from Server")
	})

	http.HandleFunc("/units", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		units := getUnits(db)
		err := json.NewEncoder(w).Encode(units)
		if err != nil {
			http.Error(w, "Failed to write to JSON", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "GET Request success (getUnits)")

	})

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
