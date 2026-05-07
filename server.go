package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
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

	http.HandleFunc("/units/status/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		// Pluck status
		status := r.URL.Path[len("/units/status/"):]

		units := getUnitsByStatus(db, status)
		err := json.NewEncoder(w).Encode(units)
		if err != nil {
			http.Error(w, "Failed to write to JSON", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "GET REQUEST success (getUnitsByStatus)")
	})

	http.HandleFunc("/units/squadron/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		//Pluck the squadron
		squadron := r.URL.Path[len("/units/squadron/"):]

		units := getUnitsBySquadron(db, squadron)
		err := json.NewEncoder(w).Encode(units)
		if err != nil {
			http.Error(w, "Failed to write to JSON", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "GET REQUEST Success (getUnitsBySquad)")
	})

	http.HandleFunc("/units/type/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		// Pluck the Type
		unitType := r.URL.Path[len("/units/type/"):]

		units := getUnitsByType(db, unitType)
		err := json.NewEncoder(w).Encode(units)
		if err != nil {
			http.Error(w, "Failed to write to JSON", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "GET Request Success (getUnitsByType)")
	})

	http.HandleFunc("/units/count", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		result := unitCount(db)
		err := json.NewEncoder(w).Encode(result)
		if err != nil {
			http.Error(w, "Failed to write to JSON", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "GET Request Success (unitCount)")
	})

	http.HandleFunc("/units/nearby", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		latStr := r.URL.Query().Get("lat")
		lonStr := r.URL.Query().Get("lon")
		radiusStr := r.URL.Query().Get("radius")

		lat, err := strconv.ParseFloat(latStr, 64)
		if err != nil {
			http.Error(w, "Failed to Parse lat", http.StatusBadRequest)
			return
		}
		lon, err := strconv.ParseFloat(lonStr, 64)
		if err != nil {
			http.Error(w, "Failed to Parse lon", http.StatusBadRequest)
			return
		}
		radius, err := strconv.ParseFloat(radiusStr, 64)
		if err != nil {
			http.Error(w, "Failed to Parse radius", http.StatusBadRequest)
			return
		}

		units := getUnits(db)
		result := []Unit{}

		for _, unit := range units {
			if haversine(lat, lon, unit.Lat, unit.Lon) <= radius {
				result = append(result, unit)
			}
		}

		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			http.Error(w, "Failed to write to JSON", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "GET Request Success (units/nearby)")
	})

	http.HandleFunc("/unit/", func(w http.ResponseWriter, r *http.Request) {
		// Pluck the ID from R
		id := r.URL.Path[len("/unit/"):]

		// IF NOT A GET METHOD RETURN
		switch r.Method {
		case http.MethodGet:
			unit := getUnitByID(db, id)
			err := json.NewEncoder(w).Encode(unit)
			if err != nil {
				http.Error(w, "Failed to write to JSON", http.StatusBadRequest)
				return
			}
			fmt.Fprintf(w, "GET Request Success (getUnitByID)")
		case http.MethodDelete:
			deleteUnit(db, id)
			fmt.Fprintf(w, "DELETE Request Success (deleteUnit)")
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}

	})

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
