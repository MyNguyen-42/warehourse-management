package model

import (
	"gitlab.ghn.vn/nhanh/backend/work-shift-service/pkg/models/enum"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserActionHistory struct {
	Id         primitive.ObjectID     `bson:"_id,omitempty"`
	HubId      string                 `bson:"hub_id"`
	LayoutCode string                 `bson:"layout_code"`

	HubAreaId     primitive.ObjectID     `bson:"hub_area_id, omitempty"`
	HubScheduleId primitive.ObjectID     `bson:"hub_schedule_id, omitempty"`
	HubShiftId    primitive.ObjectID     `bson:"hub_shift_id, omitempty"`
	Action     enum.UserActionHistory `bson:"action"`
	CreatedAt  time.Time              `bson:"created_at"`
	CreatedBy  string                 `bson:"created_by"`
}
