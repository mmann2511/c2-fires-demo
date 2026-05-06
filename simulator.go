package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var units = []Unit{
	{ID: "FURY-01", Type: "TANK", Squadron: "3rd Armored Division", Lat: 31.5497, Lon: -97.1469, Status: "ACTIVE", TimeStamp: time.Now().UTC().Format(time.RFC3339)},
	{ID: "FURY-02", Type: "TANK", Squadron: "3rd Armored Division", Lat: 31.5512, Lon: -97.1502, Status: "ACTIVE", TimeStamp: time.Now().UTC().Format(time.RFC3339)},
	{ID: "EAGLE-01", Type: "F-16", Squadron: "20th Fighter Wing", Lat: 34.0489, Lon: -80.9765, Status: "AIRBORNE", TimeStamp: time.Now().UTC().Format(time.RFC3339)},
	{ID: "EAGLE-02", Type: "F-16", Squadron: "20th Fighter Wing", Lat: 34.0512, Lon: -80.9801, Status: "AIRBORNE", TimeStamp: time.Now().UTC().Format(time.RFC3339)},
	{ID: "VIPER-01", Type: "AH-64 APACHE", Squadron: "101st Airborne", Lat: 36.5354, Lon: -87.3594, Status: "AIRBORNE", TimeStamp: time.Now().UTC().Format(time.RFC3339)},
	{ID: "VIPER-02", Type: "AH-64 APACHE", Squadron: "101st Airborne", Lat: 36.5371, Lon: -87.3612, Status: "AIRBORNE", TimeStamp: time.Now().UTC().Format(time.RFC3339)},
	{ID: "BRAVO-01", Type: "INFANTRY", Squadron: "82nd Airborne", Lat: 35.1395, Lon: -79.0146, Status: "ACTIVE", TimeStamp: time.Now().UTC().Format(time.RFC3339)},
	{ID: "BRAVO-02", Type: "INFANTRY", Squadron: "82nd Airborne", Lat: 35.1412, Lon: -79.0163, Status: "ACTIVE", TimeStamp: time.Now().UTC().Format(time.RFC3339)},
	{ID: "GHOST-01", Type: "RECON", Squadron: "75th Ranger Regiment", Lat: 31.3543, Lon: -110.9523, Status: "ACTIVE", TimeStamp: time.Now().UTC().Format(time.RFC3339)},
	{ID: "GHOST-02", Type: "RECON", Squadron: "75th Ranger Regiment", Lat: 31.3561, Lon: -110.9541, Status: "ACTIVE", TimeStamp: time.Now().UTC().Format(time.RFC3339)},
}

func runSimulator(serverURL string) {
	for _, unit := range units {
		go func(unit Unit) { // launches all at same time
			for {
				unit.Lat += (rand.Float64() - 0.5) * 0.01
				unit.Lon += (rand.Float64() - 0.5) * 0.01
				sendUpdate(serverURL, unit)
				time.Sleep(time.Duration(5) * time.Second)
			}
		}(unit)
	}
}

func sendUpdate(serverURL string, unit Unit) {
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
