package common

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Version struct {
	Id        primitive.ObjectID `bson:"_id"`
	Version   string             `bson:"version"`
	FileId    string             `bson:"file_id"`
	CreatedAt time.Time          `bson:"created_at"`
	CreatedBy string             `bson:"created_by"`
}
