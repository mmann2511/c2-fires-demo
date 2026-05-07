# unit-tracker-go

A small-scale Command and Control (C2) data ingestion pipeline built in Go.

## Overview
A Go HTTP server receives concurrent unit position and status updates from a simulator, 
writing to a PostgreSQL database using REST API endpoints. Demonstrates concurrent data 
ingestion using goroutines and a relational database backend.

## Stack
- Go
- PostgreSQL
- REST API (net/http)
- Goroutines for concurrency

## Endpoints
- `POST /unit` — insert or update a unit
- `GET /units` — return all units
- `GET /unit/{id}` — return a specific unit
- `GET /units/status/{status}` — filter by status
- `GET /units/type/{type}` — filter by unit type
- `GET /units/squadron/{squadron}` — filter by squadron
- `GET /units/count` — total unit count
- `GET /units/nearby?lat={lat}&lon={lon}&radius={radius}` — units within radius (miles)
- `DELETE /unit/{id}` — remove a unit

## Setup
Set the database password as an environment variable:
```
DB_PASSWORD=your_password_here
```