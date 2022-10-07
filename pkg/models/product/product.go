package product

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Product struct {
	ObjectId  primitive.ObjectID `bson:"_id,omitempty"`
	ProductId string             `bson:"product_id,omitempty"`
	Name      string             `bson:"name,omitempty"`
	Price     int                `bson:"price,omitempty"`
	Quantity  int                `bson:"quantity"`
	CreatedAt time.Time          `bson:"created_at"`
	CreatedBy string             `bson:"created_by"`
}
