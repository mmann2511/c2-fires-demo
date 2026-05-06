package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Unit struct {
	ID        string
	Type      string
	Squadron  string
	Lat       float64
	Lon       float64
	Status    string
	TimeStamp string
}

func insertUnit(db *sql.DB, unit Unit) {
	_, err := db.Exec(`INSERT OR REPLACE INTO units (
					unit_id, unit_type, squadron, lat, lon, status, time_stamp)
					VALUES ($1, $2, $3, $4, $5, $6, $7);`,
		unit.ID, unit.Type, unit.Squadron, unit.Lat, unit.Lon, unit.Status, unit.TimeStamp)
	if err != nil {
		log.Fatal("Failed insertUnit:", err)
	}

	fmt.Println("Success insertUnit")
}

func getUnits(db *sql.DB) []Unit {
	units := []Unit{}

	// now i need to make connecttion to db
	rows, err := db.Query("SELECT unit_id, unit_type, squadron, lat, lon, status, time_stamp FROM units")
	if err != nil {
		log.Fatal("Failed getUnits Query:", err)
	}
	defer rows.Close() // remeber to close after

	for rows.Next() {
		var unit Unit
		err := rows.Scan(&unit.ID, &unit.Type, &unit.Squadron, &unit.Lat, &unit.Lon, &unit.Status, &unit.TimeStamp)
		if err != nil {
			log.Fatal("getUnits rows.Scan failed:", err)
		}
		units = append(units, unit)
	}

	fmt.Println("getUnits success")

	return units

}

func getUnitByID(db *sql.DB, unitID string) Unit {
	var unit Unit
	err := db.QueryRow("SELECT unit_id, unit_type, squadron, lat, lon, status, time_stamp FROM units WHERE unit_id = $1", unitID).Scan(&unit.ID, &unit.Type, &unit.Squadron, &unit.Lat, &unit.Lon, &unit.Status, &unit.TimeStamp)
	if err != nil {
		log.Fatal("Failed getUnitByID:", err)
	}

	return unit
}

func deleteUnit(db *sql.DB, unitID string) {
	_, err := db.Exec("DELETE FROM units WHERE unit_id = $1", unitID)
	if err != nil {
		log.Fatal("Failed deleteUnit:", err)
	}
}
