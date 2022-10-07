package sorting

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type PackageForward struct {
	Id           primitive.ObjectID `bson:"_id"`
	Version      string             `bson:"version"`
	FromRegion   string             `bson:"from_region"`
	FromHubId    string             `bson:"from_hub_id"`
	ToRegion     string             `bson:"to_region"`
	ToHubId      string             `bson:"to_hub_id"`
	ForwardHubId string             `bson:"forward_hub_id"`
	CreatedAt    time.Time          `bson:"created_at"`
	CreatedBy    string             `bson:"created_by"`
}
