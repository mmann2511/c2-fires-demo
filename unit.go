package main

import (
	"database/sql"
	"fmt"
	"log"
	"math"

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

func getUnitsByStatus(db *sql.DB, status string) []Unit {
	units := []Unit{}

	rows, err := db.Query("SELECT unit_id, unit_type, squadron, lat, lon, status, time_stamp FROM units WHERE status = $1", status)
	if err != nil {
		log.Fatal("Failed getUnitsByStatus Query:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var unit Unit
		err := rows.Scan(&unit.ID, &unit.Type, &unit.Squadron, &unit.Lat, &unit.Lon, &unit.Status, &unit.TimeStamp)
		if err != nil {
			log.Fatal("Failed rows.Scan unitByStatus:", err)
		}
		units = append(units, unit)
	}
	fmt.Println("getUnitsByStatus Success")

	return units
}

func getUnitsBySquadron(db *sql.DB, squadron string) []Unit {
	units := []Unit{}

	// make a query
	rows, err := db.Query("SELECT unit_id, unit_type, squadron, lat, lon, status, time_stamp FROM units WHERE squadron = $1", squadron)
	if err != nil {
		log.Fatal("Failed getUnitsBySquad query:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var unit Unit
		err := rows.Scan(&unit.ID, &unit.Type, &unit.Squadron, &unit.Lat, &unit.Lon, &unit.Status, &unit.TimeStamp)
		if err != nil {
			log.Fatal("Failed rows.Scan getUnitsBySquad:", err)
		}
		units = append(units, unit)
	}

	fmt.Println("getUnitsBySquadron Success")

	return units
}

func getUnitsByType(db *sql.DB, unitType string) []Unit {
	units := []Unit{}

	// make a query
	rows, err := db.Query("SELECT unit_id, unit_type, squadron, lat, lon, status, time_stamp FROM units WHERE unit_type = $1", unitType)
	if err != nil {
		log.Fatal("Failed getUnitsByType query", err)
	}
	defer rows.Close()

	for rows.Next() {
		var unit Unit
		err := rows.Scan(&unit.ID, &unit.Type, &unit.Squadron, &unit.Lat, &unit.Lon, &unit.Status, &unit.TimeStamp)
		if err != nil {
			log.Fatal("Failed rows.Scan unitsByType:", err)
		}
		units = append(units, unit)
	}
	fmt.Println("getUnitsByType Success")

	return units
}

func unitCount(db *sql.DB) int {
	var result int
	err := db.QueryRow("SELECT COUNT(*) FROM units").Scan(&result)
	if err != nil {
		log.Fatal("Failed unitCount query:", err)
	}

	fmt.Println("unitCount success")

	return result

}

func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 3958.8
	lat1 = lat1 * math.Pi / 180
	lon1 = lon1 * math.Pi / 180
	lat2 = lat2 * math.Pi / 180
	lon2 = lon2 * math.Pi / 180

	dLat := lat2 - lat1
	dLon := lon2 - lon1

	a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(lat1)*math.Cos(lat2)*math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c
}

func deleteUnit(db *sql.DB, unitID string) {
	_, err := db.Exec("DELETE FROM units WHERE unit_id = $1", unitID)
	if err != nil {
		log.Fatal("Failed deleteUnit:", err)
	}
}
