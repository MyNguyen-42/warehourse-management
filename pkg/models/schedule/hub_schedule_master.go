package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type HubScheduleMaster struct {
	Id           primitive.ObjectID `bson:"_id"`
	HubId        string             `bson:"hub_id"`
	ScheduleId   string             `bson:"schedule_id"`
	ScheduleName string             `bson:"schedule_name"`
	Weekday      []time.Weekday     `bson:"weekday"`
	ValidAt      time.Time          `bson:"valid_at"`
	ExpiredAt    time.Time          `bson:"expired_at"`
	Shifts       []*ScheduleShift   `bson:"shifts"`
	EarlyIn      int32              `bson:"early_in"`
	LateIn       int32              `bson:"late_in"`
	Active       bool               `bson:"active"`
	CreatedAt    time.Time          `bson:"created_at"`
	CreatedBy    int32              `bson:"created_by"`
	UpdatedAt    time.Time          `bson:"updated_at"`
	UpdatedBy    int32              `bson:"updated_by"`
}

type ScheduleShift struct {
	Id    int32  `bson:"id"`
	Name  string `bson:"name"`
	Start string `bson:"start"`
	End   string `bson:"end"`
}

type UpdateSchedule struct {
	ScheduleName string           `bson:"schedule_name,omitempty"`
	RepeatType   int32            `bson:"repeat_type"`
	Weekday      []time.Weekday   `bson:"weekday"`
	Shifts       []*ScheduleShift `bson:"shifts"`
	EarlyIn      int32            `bson:"early_in"`
	LateIn       int32            `bson:"late_in"`
	UpdatedAt    time.Time        `bson:"updated_at"`
	UpdatedBy    int32            `bson:"updated_by"`
}
