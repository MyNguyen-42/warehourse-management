package models

import "go.mongodb.org/mongo-driver/mongo"

type (
	Connection struct {
		ProductCollection   *mongo.Collection
		WarehouseCollection *mongo.Collection
		OrdersCollection    *mongo.Collection
	}
)
