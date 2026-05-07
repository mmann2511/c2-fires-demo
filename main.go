package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	connStr := fmt.Sprintf("host=127.0.0.1 port=5432 user=postgres password=%s dbname=unit_tracker sslmode=disable",
		os.Getenv("DB_PASSWORD"))

	db, err := sql.Open("postgres", connStr)
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
