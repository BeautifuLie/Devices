package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id"`
	DeviceID  string             `bson:"deviceId" json:"deviceId"`
	StartDate int                `bson:"startDate" json:"startDate"`
	EndDate   int                `bson:"endDate" json:"endDate"`
}
