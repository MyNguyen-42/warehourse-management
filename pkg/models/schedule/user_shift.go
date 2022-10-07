package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserShift struct {
	Id          primitive.ObjectID `bson:"_id"`
	AreaShiftId primitive.ObjectID `bson:"area_shift_id"`
	HubId       string             `bson:"hub_id"`
	AreaCode    string             `bson:"area_code"`
	ShiftId     string             `bson:"shift_id"`
	UserId      int32              `bson:"user_id"`
	StartHour   string             `bson:"start_hour"`
	EndHour     string             `bson:"end_hour"`
	CreatedAt   time.Time          `bson:"created_at"`
	CreatedBy   int32              `bson:"created_by"`
	UpdatedAt   time.Time          `bson:"updated_at"`
	UpdatedBy   int32              `bson:"updated_by"`
}
