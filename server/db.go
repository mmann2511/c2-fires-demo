package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var units = []Unit{
	{ID: "TACP-1", Type: "OPERATOR", Squadron: "7th_ASOS", Lat: 31.8457, Lon: -106.4309, Status: "ACTIVE", TimeStamp: time.Now().UTC().Format(time.RFC3339)},
	{ID: "TACP-2", Type: "OPERATOR", Squadron: "7th_ASOS", Lat: 31.8521, Lon: -106.4401, Status: "ACTIVE", TimeStamp: time.Now().UTC().Format(time.RFC3339)},
	{ID: "FURY-1", Type: "TANK", Squadron: "1st_Armored", Lat: 31.7823, Lon: -106.3201, Status: "ACTIVE", TimeStamp: time.Now().UTC().Format(time.RFC3339)},
	{ID: "FURY-2", Type: "TANK", Squadron: "1st_Armored", Lat: 31.7901, Lon: -106.3389, Status: "ACTIVE", TimeStamp: time.Now().UTC().Format(time.RFC3339)},
	{ID: "FURY-3", Type: "TANK", Squadron: "1st_Armored", Lat: 31.7765, Lon: -106.3102, Status: "ACTIVE", TimeStamp: time.Now().UTC().Format(time.RFC3339)},
	{ID: "GHOST-1", Type: "RECON", Squadron: "75th_Rangers", Lat: 31.9102, Lon: -106.5001, Status: "ACTIVE", TimeStamp: time.Now().UTC().Format(time.RFC3339)},
	{ID: "GHOST-2", Type: "RECON", Squadron: "75th_Rangers", Lat: 31.9234, Lon: -106.5123, Status: "ACTIVE", TimeStamp: time.Now().UTC().Format(time.RFC3339)},
	{ID: "GHOST-3", Type: "RECON", Squadron: "75th_Rangers", Lat: 31.8998, Lon: -106.4889, Status: "ACTIVE", TimeStamp: time.Now().UTC().Format(time.RFC3339)},
	{ID: "BRAVO-1", Type: "INFANTRY", Squadron: "82nd_Airborne", Lat: 31.8201, Lon: -106.3701, Status: "ACTIVE", TimeStamp: time.Now().UTC().Format(time.RFC3339)},
	{ID: "BRAVO-2", Type: "INFANTRY", Squadron: "82nd_Airborne", Lat: 31.8312, Lon: -106.3812, Status: "ACTIVE", TimeStamp: time.Now().UTC().Format(time.RFC3339)},
	{ID: "BRAVO-3", Type: "INFANTRY", Squadron: "82nd_Airborne", Lat: 31.8089, Lon: -106.3612, Status: "ACTIVE", TimeStamp: time.Now().UTC().Format(time.RFC3339)},
	{ID: "VIPER-1", Type: "AH-64", Squadron: "101st_Airborne", Lat: 31.8701, Lon: -106.4501, Status: "ACTIVE", TimeStamp: time.Now().UTC().Format(time.RFC3339)},
	{ID: "VIPER-2", Type: "AH-64", Squadron: "101st_Airborne", Lat: 31.8812, Lon: -106.4612, Status: "ACTIVE", TimeStamp: time.Now().UTC().Format(time.RFC3339)},
	{ID: "EAGLE-1", Type: "F-16", Squadron: "20th_Fighter_Wing", Lat: 31.9401, Lon: -106.5301, Status: "AIRBORNE", TimeStamp: time.Now().UTC().Format(time.RFC3339)},
	{ID: "EAGLE-2", Type: "F-16", Squadron: "20th_Fighter_Wing", Lat: 31.9512, Lon: -106.5412, Status: "AIRBORNE", TimeStamp: time.Now().UTC().Format(time.RFC3339)},
	{ID: "WOLF-1", Type: "RECON", Squadron: "160th_SOAR", Lat: 31.7601, Lon: -106.2901, Status: "ACTIVE", TimeStamp: time.Now().UTC().Format(time.RFC3339)},
	{ID: "WOLF-2", Type: "RECON", Squadron: "160th_SOAR", Lat: 31.7712, Lon: -106.3012, Status: "ACTIVE", TimeStamp: time.Now().UTC().Format(time.RFC3339)},
	{ID: "STRIKER-1", Type: "ARTILLERY", Squadron: "75th_FA", Lat: 31.7401, Lon: -106.2501, Status: "ACTIVE", TimeStamp: time.Now().UTC().Format(time.RFC3339)},
	{ID: "STRIKER-2", Type: "ARTILLERY", Squadron: "75th_FA", Lat: 31.7512, Lon: -106.2612, Status: "ACTIVE", TimeStamp: time.Now().UTC().Format(time.RFC3339)},
	{ID: "STRIKER-3", Type: "ARTILLERY", Squadron: "75th_FA", Lat: 31.7289, Lon: -106.2389, Status: "ACTIVE", TimeStamp: time.Now().UTC().Format(time.RFC3339)},
}

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
	seedDB(db)

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

func seedDB(db *sql.DB) {
	var count int
	db.QueryRow("SELECT COUNT(*) FROM units").Scan(&count)
	if count == 0 {
		for _, unit := range units {
			insertUnit(db, unit)
		}
		fmt.Println("Successfully Seeded units table")
	}
}
