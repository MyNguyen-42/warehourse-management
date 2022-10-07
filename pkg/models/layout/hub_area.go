package model

import (
	"gitlab.ghn.vn/nhanh/backend/work-shift-service/interface/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type HubArea struct {
	Id            primitive.ObjectID   `bson:"_id"`
	HubId         string               `bson:"hub_id"`
	LayoutCode    string               `bson:"layout_code"`
	AreaCode      string               `bson:"area_code"`
	AreaName      string               `bson:"area_name"`
	Description   string               `bson:"description"`
	Actions       []string             `bson:"actions"`
	HubScheduleId primitive.ObjectID   `bson:"hub_schedule_id, omitempty"`
	HubSchedules  []*services.Schedule `bson:"hub_schedules"`
	CheckinCode   string               `bson:"checkin_code"`
	Active        bool                 `bson:"active"`
	Tag           string               `json:"tag"`
	CreatedAt     time.Time            `bson:"created_at"`
	CreatedBy     string               `bson:"created_by"`
	UpdatedAt     time.Time            `bson:"updated_at"`
	UpdatedBy     string               `bson:"updated_by"`
}

type UpdateArea struct {
	AreaName      string               `bson:"area_name,omitempty"`
	HubScheduleId primitive.ObjectID   `bson:"hub_schedule_id,omitempty"`
	Actions       *[]string            `bson:"actions,omitempty"`
	Active        *bool                `bson:"active,omitempty"`
	UpdatedAt     time.Time            `bson:"updated_at,omitempty"`
	UpdatedBy     string               `bson:"updated_by,omitempty"`
	Description   string               `bson:"description,omitempty"`
	HubSchedules  *[]*services.Schedule `bson:"hub_schedules,omitempty"`
	CheckinCode   string               `bson:"checkin_code,omitempty"`
}
