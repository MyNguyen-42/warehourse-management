package common

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ConfigVersion struct {
	Id        primitive.ObjectID `bson:"_id"`
	ConfigId  string             `bson:"config_id"`
	Version   string             `bson:"version"`
	CreatedAt time.Time          `bson:"created_at"`
	CreatedBy string             `bson:"created_by"`
}
