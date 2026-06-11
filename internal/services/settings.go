package services

import (
	"context"
	"time"

	"attendance-backend/internal/config"
	"attendance-backend/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var defaultOffice = models.OfficeSettings{
	Name:         "Main Office",
	Latitude:     13.7563,
	Longitude:    100.5018,
	RadiusMeters: 200.0,
}

func GetOfficeSettings() (models.OfficeSettings, error) {
	collection := config.DB.Collection("office_settings")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var settings models.OfficeSettings
	err := collection.FindOne(ctx, bson.M{}).Decode(&settings)
	if err == mongo.ErrNoDocuments {
		return defaultOffice, nil
	}
	if err != nil {
		return defaultOffice, err
	}
	return settings, nil
}

func SaveOfficeSettings(settings models.OfficeSettings) error {
	collection := config.DB.Collection("office_settings")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	opts := options.Replace().SetUpsert(true)
	_, err := collection.ReplaceOne(ctx, bson.M{}, settings, opts)
	return err
}
