# c2-fires-demo

A small-scale Command and Control (C2) fires deconfliction system demonstrating a C++/Go/PostgreSQL architecture.

## Overview
A ground unit simulator moves toward a target, calls in the target location via REST API, then moves to a safe distance. A C++ artillery system polls for active targets, checks for nearby friendly units, and fires only when the area is clear. All data flows through a Go HTTP server backed by PostgreSQL.

## Stack
- Go — HTTP server, REST API, TACP simulator
- C++ (libcurl) — artillery deconfliction system
- PostgreSQL — operational data store
- Goroutines — concurrent unit position updates

## Architecture
C++ Artillery → GET /targets, GET /units/nearby → Go Server → PostgreSQL
Go Simulator → POST /unit, POST /target → Go Server → PostgreSQL

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
- `POST /target` — report a new target
- `GET /targets` — return all targets
- `GET /target/{id}` — return a specific target
- `PUT /target/{id}?status={status}` — update target status

## Setup
Set the database password as an environment variable:
DB_PASSWORD=your_password_here
