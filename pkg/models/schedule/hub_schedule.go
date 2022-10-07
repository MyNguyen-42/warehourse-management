package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type HubSchedule struct {
	Id          primitive.ObjectID `bson:"_id"`
	DisplayCode string             `bson:"display_code"`
	HubId       string             `bson:"hub_id"`
	HubAreaId   primitive.ObjectID `bson:"hub_area_id, omitempty"`
	CreatedAt   time.Time          `bson:"created_at"`
	CreatedBy   string             `bson:"created_by"`
	UpdatedAt   time.Time          `bson:"updated_at"`
	UpdatedBy   string             `bson:"updated_by"`
}
