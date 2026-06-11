package models

type OfficeSettings struct {
	Name         string  `bson:"name"          json:"name"`
	Latitude     float64 `bson:"latitude"      json:"latitude"`
	Longitude    float64 `bson:"longitude"     json:"longitude"`
	RadiusMeters float64 `bson:"radius_meters" json:"radius_meters"`
}
