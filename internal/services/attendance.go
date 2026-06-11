package services

import (
	"context"
	"errors"
	"math"
	"time"

	"attendance-backend/internal/config"
	"attendance-backend/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	officeLatitude  = 13.7563
	officeLongitude = 100.5018
	maxRadiusMeters = 200.0
)

var allowedSessions = map[string]bool{
	"morning":   true,
	"lunch":     true,
	"afternoon": true,
	"evening":   true,
}

// haversine returns the distance in meters between two lat/lon points.
func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const earthRadius = 6371000.0

	φ1 := lat1 * math.Pi / 180
	φ2 := lat2 * math.Pi / 180
	Δφ := (lat2 - lat1) * math.Pi / 180
	Δλ := (lon2 - lon1) * math.Pi / 180

	a := math.Sin(Δφ/2)*math.Sin(Δφ/2) +
		math.Cos(φ1)*math.Cos(φ2)*math.Sin(Δλ/2)*math.Sin(Δλ/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadius * c
}

func ClockIn(req models.ClockInRequest) (*models.AttendanceLog, error) {
	if !allowedSessions[req.Session] {
		return nil, errors.New("invalid session: must be morning, lunch, afternoon, or evening")
	}

	distance := haversine(req.Latitude, req.Longitude, officeLatitude, officeLongitude)

	if distance > maxRadiusMeters {
		return nil, errors.New("Access denied: You are outside the 200-meter radius.")
	}

	entry := &models.AttendanceLog{
		EmployeeID:     req.EmployeeID,
		Session:        req.Session,
		Latitude:       req.Latitude,
		Longitude:      req.Longitude,
		ImageBase64:    req.ImageBase64,
		Timestamp:      time.Now().UTC(),
		DistanceMeters: math.Round(distance*100) / 100,
	}

	collection := config.DB.Collection("attendance_logs")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, entry)
	if err != nil {
		return nil, err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		entry.ID = oid
	}

	return entry, nil
}

func GetLogs() ([]models.AttendanceLog, error) {
	collection := config.DB.Collection("attendance_logs")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var logs []models.AttendanceLog
	if err := cursor.All(ctx, &logs); err != nil {
		return nil, err
	}

	return logs, nil
}
