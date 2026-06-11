package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ClockInRequest struct {
	EmployeeID  string  `json:"employee_id"`
	Session     string  `json:"session"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	ImageBase64 string  `json:"image_base64"`
}

type AttendanceLog struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	EmployeeID     string             `bson:"employee_id"   json:"employee_id"`
	Session        string             `bson:"session"       json:"session"`
	Latitude       float64            `bson:"latitude"      json:"latitude"`
	Longitude      float64            `bson:"longitude"     json:"longitude"`
	ImageBase64    string             `bson:"image_base64"  json:"image_base64"`
	Timestamp      time.Time          `bson:"timestamp"     json:"timestamp"`
	DistanceMeters float64            `bson:"distance_meters" json:"distance_meters"`
}
