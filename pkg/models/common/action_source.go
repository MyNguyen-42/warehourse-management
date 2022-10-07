package common

import "time"

type (
	ActionSource struct {
		Id     string    `json:"id" bson:"id"`
		Time   time.Time `json:"time" bson:"time"`
		Source string    `json:"source" bson:"source"`
	}
)
