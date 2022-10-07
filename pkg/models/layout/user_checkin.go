package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserCheckin struct {
	Id          primitive.ObjectID `bson:"_id"`
	UserId      string             `bson:"user_id"`
	HubId       string             `bson:"hub_id"`
	HubAreaId   primitive.ObjectID `bson:"hub_area_id"`
	HubShitId   primitive.ObjectID `bson:"hub_shift_id"`
	CheckAt     time.Time          `bson:"check_at"`
	CreatedAt   time.Time          `bson:"created_at"`
	CreatedBy   string             `bson:"created_by"`
	UpdatedAt   time.Time          `bson:"updated_at"`
	UpdatedBy   string             `bson:"updated_by"`
	CheckinType int32              `bson:"checkin_type"`
}
