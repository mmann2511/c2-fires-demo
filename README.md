# c2-fires-demo

A small-scale Command and Control (C2) data ingestion pipeline demonstrating a C++/Go/PostgreSQL architecture with concurrent field unit simulation, atomic transactions, and a real-time web dashboard.

## Overview
20 C++ field units run as concurrent threads, continuously sending position updates to a Go HTTP server via REST API. Units randomly report threats, triggering atomic transactions that simultaneously update unit status and insert threat records. A C++ artillery consumer polls for active threats, checks for nearby friendly units using the Haversine formula, and approves fire missions when the area is clear. A real-time HTML dashboard displays the live operational picture.

## Stack
- **Go** — HTTP server, REST API, PostgreSQL interface, real-time dashboard serving
- **C++ (libcurl, nlohmann/json)** — concurrent field unit simulator, artillery fire control consumer
- **PostgreSQL** — operational data store
- **HTML/JavaScript** — real-time command dashboard

## Architecture
- C++ Field Units (20 threads) → PUT /unit → Go Server → PostgreSQL
- C++ Field Units → POST /report-threat (atomic transaction) → Go Server → PostgreSQL
- C++ Artillery Consumer → GET /targets, GET /units/nearby → Go Server → PostgreSQL
- Browser Dashboard → GET /units, GET /targets → Go Server → PostgreSQL

## Key Concepts Demonstrated
- Concurrent C++ clients hitting Go simultaneously via goroutines
- Atomic transactions — threat INSERT and unit status UPDATE succeed or fail together
- Haversine formula for GPS-based deconfliction
- REST API design with Go's net/http
- Real-time dashboard consuming live operational data

## Field Operations (hardware/)
- `ground_unit.cpp` — single TACP ground unit that moves toward a target and reports threats
- `main.cpp` — C++ artillery fire control consumer

## Endpoints
- `POST /unit` — insert a unit
- `PUT /unit/{id}` — update unit position and status
- `GET /units` — return all units
- `GET /unit/{id}` — return a specific unit
- `GET /units/status/{status}` — filter by status
- `GET /units/type/{type}` — filter by unit type
- `GET /units/squadron/{squadron}` — filter by squadron
- `GET /units/count` — total unit count
- `GET /units/nearby?lat={lat}&lon={lon}&radius={radius}` — units within radius (miles)
- `DELETE /unit/{id}` — remove a unit
- `POST /target` — report a new target
- `POST /report-threat` — atomic transaction: insert threat + update unit status
- `GET /targets` — return all targets
- `GET /target/{id}` — return a specific target
- `PUT /target/{id}?status={status}` — update target status
- `GET /dashboard` — real-time operational dashboard

## Setup
Set the database password as an environment variable:
DB_PASSWORD=your_password_here
