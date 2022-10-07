package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserScheduleConfig struct {
	Id         primitive.ObjectID `bson:"_id"`
	HubId      string             `bson:"hub_id"`
	ScheduleId string             `bson:"schedule_id"`
	ShiftId    int32              `bson:"shift_id"`
	UserId     string             `bson:"user_id"`
	UserName   string             `bson:"user_name"`
	CreatedAt  time.Time          `bson:"created_at"`
	CreatedBy  string             `bson:"created_by"`
	UpdatedAt  time.Time          `bson:"updated_at"`
	UpdatedBy  string             `bson:"updated_by"`
}

type UpdateUserShift struct {
	ShiftId   int32     `bson:"shift_id"`
	UpdatedAt time.Time `bson:"updated_at"`
	UpdatedBy string    `bson:"updated_by"`
}
