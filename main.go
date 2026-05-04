package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "units.db")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	fmt.Println("Connected to database")

	setupDB(db)
	go startServer(db)
	time.Sleep(time.Second)
	runSimulator("http://localhost:8080")
	select {}

}
