package soundConfig

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type RouteConfigDataDto struct {
	Id         primitive.ObjectID `bson:"_id"`
	HubId      string             `bson:"hub_id"`
	Action     string             `bson:"action"`
	Layout     string             `bson:"layout"`
	FileName   string             `bson:"file_name"`
	RouteCode  string             `bson:"route_code"`
	Condition  int32              `bson:"condition"`
	NextHubIds []string           `bson:"next_hub_ids"`
	FileId     string             `bson:"file_id"`
	CreatedAt  time.Time          `bson:"created_at"`
	CreatedBy  string             `bson:"created_by"`
}

type RouteConfigDataRaw struct {
	Action    string `json:"action"`
	ToHubId   string `json:"to_hub_id"`
	ToHubName string `json:"to_hub_name"`
	RouteCode string `json:"route_code"`
	Condition string `json:"condition"`
	FileName  string `json:"file_name"`
	RowNumber int32
}
