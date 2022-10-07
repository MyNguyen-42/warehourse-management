package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserSchedule struct {
	Id            primitive.ObjectID `bson:"_id"`
	HubId         string             `bson:"hub_id"`
	ScheduleId    string             `bson:"schedule_id"`
	HubScheduleId primitive.ObjectID `bson:"hub_schedule_id"`
	ShiftId       int32              `bson:"shift_id"`
	UserId        string             `bson:"user_id"`
	UserName      string             `bson:"user_name"`
	InTime        time.Time          `bson:"in_time"`
	OutTime       time.Time          `bson:"out_time"`
	CheckedInAt   *time.Time         `bson:"checked_in_at"`
	CheckedOutAt  *time.Time         `bson:"checked_out_at"`
	CreatedAt     time.Time          `bson:"created_at"`
	CreatedBy     string             `bson:"created_by"`
	UpdatedAt     time.Time          `bson:"updated_at"`
	UpdatedBy     string             `bson:"updated_by"`
}
