package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id"`
	DeviceID  string             `bson:"deviceId" json:"deviceId"`
	StartDate time.Time          `bson:"startDate" json:"startDate"`
	EndDate   time.Time          `bson:"endDate" json:"endDate"`
}
