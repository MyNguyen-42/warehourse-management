package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type HubLayout struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	HubId     string             `bson:"hub_id"`
	Code      string             `bson:"code"`
	Name      string             `bson:"name"`
	Deleted   bool               `bson:"deleted"`
	CreatedAt time.Time          `bson:"created_at"`
	CreatedBy string             `bson:"created_by"`
	UpdatedAt time.Time          `bson:"updated_at"`
	UpdatedBy string             `bson:"updated_by"`
}
