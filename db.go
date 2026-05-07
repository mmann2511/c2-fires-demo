package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func setupDB(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS units (
	unit_id TEXT PRIMARY KEY,
	unit_type TEXT,
	squadron TEXT,
	lat REAL, 
	lon REAL,
	status TEXT,
	time_stamp TEXT
	);`)
	if err != nil {
		log.Fatal("Failed to create accounts table:", err)
	}

	fmt.Println("Units table created successfully")
}
