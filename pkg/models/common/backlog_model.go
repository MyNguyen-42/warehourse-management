package common

import "time"

type BacklogModel struct {
	Key           string    `bson:"key"`
	Value         string    `bson:"value"`
	Error         string    `bson:"error"`
	Partition     int       `bson:"partition"`
	HighWaterMark int64     `bson:"high_water_mark"`
	Time          time.Time `bson:"time"`
}
