package main

import (
	"database/sql"
	"fmt"
	"log"
)

type Target struct {
	ID          string
	Description string
	Lat         float64
	Lon         float64
	Status      string
	TimeStamp   string
}

func insertTarget(db *sql.DB, target Target) {
	_, err := db.Exec(`INSERT INTO targets (
				target_id, description, lat, lon, status, time_stamp)
				VALUES($1, $2, $3, $4, $5, $6)
				ON CONFLICT (target_id) DO NOTHING`,
		target.ID, target.Description, target.Lat, target.Lon, target.Status, target.TimeStamp)
	if err != nil {
		log.Fatal("Failed to insert target:", err)
	}
	fmt.Println("Insert Target Success!")
}

func getTarget(db *sql.DB, targetID string) Target {
	var target Target
	err := db.QueryRow(`SELECT target_ID, description, lat, lon, status, time_stamp 
						FROM targets WHERE target_id = $1`, targetID).Scan(
		&target.ID, &target.Description, &target.Lat, &target.Lon, &target.Status, &target.TimeStamp)
	if err != nil {
		log.Fatal("Failed to getTarget"+target.ID, err)
	}
	fmt.Println("Success getTarget!")
	return target
}

func getTargets(db *sql.DB) []Target {
	targets := []Target{}

	// make a query
	rows, err := db.Query("SELECT target_id, description, lat, lon, status, time_stamp FROM targets")
	if err != nil {
		log.Fatal("Failed to query rows (getTargets):", err)
	}
	defer rows.Close()

	for rows.Next() {
		var target Target
		err := rows.Scan(&target.ID, &target.Description, &target.Lat, &target.Lon, &target.Status, &target.TimeStamp)
		if err != nil {
			log.Fatal("Failed to scan rows (getTargets):", err)
		}
		targets = append(targets, target)
	}

	return targets
}

func updateTargetStatus(db *sql.DB, targetID string, newStatus string) error {
	_, err := db.Exec("UPDATE targets SET status = $1 WHERE target_id = $2", newStatus, targetID)
	if err != nil {
		return err
	}

	fmt.Println("Success updateTargetStatus")
	return nil
}
