package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserSession struct {
	Id              primitive.ObjectID `bson:"_id"`
	GroupId         string             `bson:"group_id,omitempty"`
	UserId          string             `bson:"user_id"`
	HubId           string             `bson:"hub_id"`
	HubAreaId       primitive.ObjectID `bson:"hub_area_id"`
	CheckinAt       time.Time          `bson:"checkin_at"`
	CreatedAt       time.Time          `bson:"created_at"`
	CreatedBy       string             `bson:"created_by"`
	UpdatedAt       time.Time          `bson:"updated_at"`
	UpdatedBy       string             `bson:"updated_by"`
	IsLeader        bool               `bson:"is_leader"`
	HubShiftId      primitive.ObjectID `bson:"hub_shift_id"`
	IsLeaderAllArea bool               `bson:"is_leader_all_area"`
}

type UpdateSessionArea struct {
	GroupId         string              `bson:"group_id,omitempty"`
	HubId           string              `bson:"hub_id,omitempty"`
	HubAreaId       primitive.ObjectID  `bson:"hub_area_id,omitempty"`
	CheckinAt       *time.Time          `bson:"checkin_at,omitempty"`
	UpdatedAt       time.Time           `bson:"updated_at"`
	UpdatedBy       string              `bson:"updated_by"`
	IsLeader        *bool               `bson:"is_leader,omitempty"`
	HubShiftId      *primitive.ObjectID `bson:"hub_shift_id,omitempty"`
	IsLeaderAllArea bool                `bson:"is_leader_all_area"`
}
