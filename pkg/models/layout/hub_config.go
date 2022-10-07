package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type HubConfig struct {
	Id          primitive.ObjectID `bson:"_id"`
	HubId       string             `bson:"hub_id"`
	Config      int32              `bson:"config"`
	Value       string             `bson:"value"`
	CreatedAt   time.Time          `bson:"created_at"`
	CreatedBy   string             `bson:"created_by"`
	UpdatedAt   time.Time          `bson:"updated_at"`
	UpdatedBy   string             `bson:"updated_by"`
	FileId      string             `bson:"file_id"`
	NeedCheckIP bool               `bson:"need_check_ip"`
}

type UpdateConfig struct {
	Value     string    `bson:"value,omitempty"`
	UpdatedAt time.Time `bson:"updated_at"`
	UpdatedBy string    `bson:"updated_by"`
	FileId    string    `bson:"file_id,omitempty"`
}
