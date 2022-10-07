package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type HubShift struct {
	Id            primitive.ObjectID `bson:"_id,omitempty"`
	HubAreaId     primitive.ObjectID `bson:"hub_area_id, omitempty"`
	HubScheduleId primitive.ObjectID `bson:"hub_schedule_id, omitempty"`
	Weekday       time.Weekday       `bson:"weekday"`
	StartHour     string             `bson:"start_hour"`
	EndHour       string             `bson:"end_hour"`
	UserIds       []string           `bson:"user_ids"`
	ForceUserIds  []string           `bson:"force_user_ids"`
	LeaderId      string             `bson:"leader_id"`
	CreatedAt     time.Time          `bson:"created_at"`
	CreatedBy     string             `bson:"created_by"`
	UpdatedAt     time.Time          `bson:"updated_at"`
	UpdatedBy     string             `bson:"updated_by"`
}

type UpdateShift struct {
	StartHour    string    `bson:"start_hour,omitempty"`
	EndHour      string    `bson:"end_hour,omitempty"`
	UserIds      []string  `bson:"user_ids"`
	UpdatedAt    time.Time `bson:"updated_at"`
	UpdatedBy    string    `bson:"updated_by"`
	LeaderId     string    `bson:"leader_id"`
	ForceUserIds []string  `bson:"force_user_ids"`
}

type HubShiftUploadRaw struct {
	Weekday      string `json:"weekday"`
	StartHour    string `json:"start_hour"`
	EndHour      string `json:"end_hour"`
	UserIds      string `json:"user_ids"`
	ForceUserIds string `json:"force_user_ids"`
	LeaderId     string `json:"leader_id"`
}
