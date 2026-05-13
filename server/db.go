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
	targetsTable(db)

	fmt.Println("Units table created successfully")
}

func targetsTable(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS targets (
	target_id TEXT PRIMARY KEY,
	description TEXT,
	lat REAL,
	lon REAL,
	status TEXT,
	time_stamp TEXT
	);`)
	if err != nil {
		log.Fatal("Failed to create targets table:", err)
	}

	fmt.Println("Targets Table created succesfully")
}
