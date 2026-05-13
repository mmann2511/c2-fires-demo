package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var units = []Unit{
	{ID: "TACP-1", Type: "OPERATOR", Squadron: "7th_ASOS", Lat: 31.8457, Lon: -106.4309, Status: "ACTIVE", TimeStamp: time.Now().UTC().Format(time.RFC3339)},
	{ID: "FURY-01", Type: "TANK", Squadron: "3rd_Armored_Division", Lat: 31.5497, Lon: -97.1469, Status: "ACTIVE", TimeStamp: time.Now().UTC().Format(time.RFC3339)},
	{ID: "FURY-02", Type: "TANK", Squadron: "3rd_Armored_Division", Lat: 31.5512, Lon: -97.1502, Status: "ACTIVE", TimeStamp: time.Now().UTC().Format(time.RFC3339)},
	{ID: "EAGLE-01", Type: "F-16", Squadron: "20th_Fighter_Wing", Lat: 34.0489, Lon: -80.9765, Status: "AIRBORNE", TimeStamp: time.Now().UTC().Format(time.RFC3339)},
	{ID: "EAGLE-02", Type: "F-16", Squadron: "20th_Fighter_Wing", Lat: 34.0512, Lon: -80.9801, Status: "AIRBORNE", TimeStamp: time.Now().UTC().Format(time.RFC3339)},
	{ID: "VIPER-01", Type: "AH-64_APACHE", Squadron: "101st_Airborne", Lat: 36.5354, Lon: -87.3594, Status: "AIRBORNE", TimeStamp: time.Now().UTC().Format(time.RFC3339)},
	{ID: "VIPER-02", Type: "AH-64_APACHE", Squadron: "101st_Airborne", Lat: 36.5371, Lon: -87.3612, Status: "AIRBORNE", TimeStamp: time.Now().UTC().Format(time.RFC3339)},
	{ID: "BRAVO-01", Type: "INFANTRY", Squadron: "82nd_Airborne", Lat: 35.1395, Lon: -79.0146, Status: "ACTIVE", TimeStamp: time.Now().UTC().Format(time.RFC3339)},
	{ID: "BRAVO-02", Type: "INFANTRY", Squadron: "82nd_Airborne", Lat: 35.1412, Lon: -79.0163, Status: "ACTIVE", TimeStamp: time.Now().UTC().Format(time.RFC3339)},
	{ID: "GHOST-01", Type: "RECON", Squadron: "75th_Ranger_Regiment", Lat: 31.3543, Lon: -110.9523, Status: "ACTIVE", TimeStamp: time.Now().UTC().Format(time.RFC3339)},
	{ID: "GHOST-02", Type: "RECON", Squadron: "75th_Ranger_Regiment", Lat: 31.3561, Lon: -110.9541, Status: "ACTIVE", TimeStamp: time.Now().UTC().Format(time.RFC3339)},
}

func runSimulator(serverURL string) {
	for _, unit := range units {
		go func(unit Unit) { // launches all at same time
			for {
				unit.Lat += (rand.Float64() - 0.5) * 0.01
				unit.Lon += (rand.Float64() - 0.5) * 0.01
				sendInsert(serverURL, unit)
				time.Sleep(time.Duration(5) * time.Second)
			}
		}(unit)
	}
}

func sendInsert(serverURL string, unit Unit) {
	body, err := json.Marshal(unit)
	if err != nil {
		log.Println("Failed to encode unit:", err)
		return
	}

	resp, err := http.Post(serverURL+"/unit", "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Println("Failed to send update:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Sent update for", unit.ID)
}

func sendTarget(serverURL string, target Target) {
	body, err := json.Marshal(target)
	if err != nil {
		log.Println("Failed to encode unit:", err)
		return
	}

	resp, err := http.Post(serverURL+"/target", "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Println("Failed to send update:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Inserted Target", target.ID)
}

func moveToward(current float64, target float64, step float64) float64 {
	if current < target {
		return current + step
	}
	return current - step
}

func tacpSimulator(serverURL string, db *sql.DB) {
	db.Exec("DELETE FROM targets WHERE target_id = 'TARGET1'")
	// reset units
	// grab unit
	tacp := getUnitByID(db, "TACP-1")
	tacp.Lat = 31.8457
	tacp.Lon = -106.4309

	targetLat := 31.9323
	targetLon := -106.5558
	startLat := tacp.Lat
	startLon := tacp.Lon

	targetReported := false

	for {
		distance := haversine(tacp.Lat, tacp.Lon, targetLat, targetLon)

		if distance > 0.4 && !targetReported {
			// move toward target
			tacp.Lat = moveToward(tacp.Lat, targetLat, 0.009)
			tacp.Lon = moveToward(tacp.Lon, targetLon, 0.009)
		} else if !targetReported {
			// within range POST TARGET
			target := Target{ID: "TARGET1", Description: "MOUNTAIN", Lat: 31.9323, Lon: -106.5558, Status: "ACTIVE", TimeStamp: time.Now().UTC().Format(time.RFC3339)}
			sendTarget(serverURL, target)
			targetReported = true
		} else {
			// move away
			tacp.Lat = moveToward(tacp.Lat, startLat, 0.005)
			tacp.Lon = moveToward(tacp.Lon, startLon, 0.005)
		}

		sendInsert(serverURL, tacp)
		fmt.Println("Distance:", distance)
		time.Sleep(5 * time.Second)
	}

}
