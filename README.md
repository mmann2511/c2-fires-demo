Building a Go HTTP server simulating a small-scale command and control data ingestion pipeline. A simulator concurrently sends unit position and status updates via REST API POST requests using goroutines. The server handles each request concurrently, writing to a SQLite database. Unit data can be retrieved via a GET endpoint returning JSON.

## Tech Stack
- Go
- SQLite
- REST API (net/http)
- Goroutines for concurrency