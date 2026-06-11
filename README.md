# Attendance Tracker — Backend API

REST API for the Attendance Tracker web application. Handles clock-in/out events with geofencing validation and face image capture.

## Tech Stack

- **Language:** Go 1.22
- **Framework:** Fiber v2
- **Database:** MongoDB Atlas
- **Task Runner:** Task (taskfile.dev)

## Project Structure

```
backend-dev/
├── cmd/
│   └── main.go                  # Entry point
├── internal/
│   ├── config/
│   │   └── database.go          # MongoDB connection
│   ├── models/
│   │   └── attendance.go        # Structs & BSON mappings
│   ├── services/
│   │   └── attendance.go        # Business logic & Haversine
│   └── controllers/
│       └── attendance.go        # HTTP handlers
├── .env.example                 # Environment variable template
├── Taskfile.yml
└── go.mod
```

## Getting Started

### Prerequisites

- [Go 1.22+](https://go.dev/dl/)
- [Task](https://taskfile.dev/installation/) — `brew install go-task`
- MongoDB Atlas account (free M0 cluster)

### Setup

1. Clone the repository and navigate to the project folder

2. Copy the environment template and fill in your values:
   ```bash
   cp .env.example .env
   ```

3. Edit `.env`:
   ```env
   MONGODB_URI=mongodb+srv://<username>:<password>@<cluster>.mongodb.net/?retryWrites=true&w=majority
   DB_NAME=AttendanceTracker
   PORT=8080
   ```

4. Run the development server:
   ```bash
   task dev
   ```
   This will automatically run `go mod tidy` before starting the server.

## API Endpoints

### POST `/api/attendance/clock-in`

Records a clock-in event. Validates geofencing (200m radius) before saving.

**Request Body:**
```json
{
  "employee_id": "EMP001",
  "session": "morning",
  "latitude": 13.7563,
  "longitude": 100.5018,
  "image_base64": "<base64-encoded-image>"
}
```

Allowed `session` values: `morning`, `lunch`, `afternoon`, `evening`

**Success Response (200):**
```json
{
  "status": "success",
  "message": "Clock-in successful",
  "distance_meters": 15.4,
  "timestamp": "2026-06-11T08:30:00Z"
}
```

**Error Response (400) — outside 200m radius:**
```json
{
  "status": "error",
  "message": "Access denied: You are outside the 200-meter radius."
}
```

---

### GET `/api/attendance/logs`

Returns all attendance records from the database.

**Response (200):**
```json
[
  {
    "id": "6849a1bc...",
    "employee_id": "EMP001",
    "session": "morning",
    "latitude": 13.7563,
    "longitude": 100.5018,
    "image_base64": "...",
    "timestamp": "2026-06-11T08:30:00Z",
    "distance_meters": 15.4
  }
]
```

## Available Tasks

| Command | Description |
|---|---|
| `task dev` | Install dependencies and run development server |
| `task build` | Build production binary (Linux/amd64) |
| `task deps` | Run `go mod tidy` only |
| `task clean` | Remove build artifacts |

## Environment Variables

| Variable | Required | Description |
|---|---|---|
| `MONGODB_URI` | Yes | MongoDB Atlas connection string |
| `DB_NAME` | No | Database name (default: `attendance_db`) |
| `PORT` | No | Server port (default: `8080`) |
