package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ConfigUserArea struct {
	Id        primitive.ObjectID `bson:"_id"`
	UserId    int32              `bson:"user_id"`
	HubId     string             `bson:"hub_id"`
	AreaCode  string             `bson:"area_code"`
	CreatedAt time.Time          `bson:"created_at"`
	CreatedBy int32              `bson:"created_by"`
	UpdatedAt time.Time          `bson:"updated_at"`
	UpdatedBy int32              `bson:"updated_by"`
}

type UpdateLocation struct {
	HubId     string    `bson:"hub_id"`
	AreaCode  string    `bson:"area_code"`
	UpdatedAt time.Time `bson:"updated_at"`
	UpdatedBy int32     `bson:"updated_by"`
}
