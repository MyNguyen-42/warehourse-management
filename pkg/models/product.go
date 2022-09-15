package models

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Product struct {
		ObjectId  primitive.ObjectID `json:"-"  bson:"_id,omitempty"`
		ProductId string             `json:"product_id,omitempty"  bson:"product_id,omitempty"`
		Name      string             `json:"name,omitempty"  bson:"name,omitempty"`
		Price     int                `json:"price,omitempty"  bson:"price,omitempty"`
		Quantity  int                `json:"quantity"  bson:"quantity"`
		Total     int                `json:"total"  bson:"total"`
	}

	InsertProductParams struct {
		ProductId string `bson:"product_id,omitempty"`
		Name      string `bson:"name,omitempty"`
		Price     int    `bson:"price,omitempty"`
	}

	UpdateProductParams struct {
		Name  string `bson:"name,omitempty"`
		Price int    `bson:"price,omitempty"`
	}

	OrderProductParams struct {
		Name     string `bson:"name,omitempty"`
		Price    int    `bson:"price,omitempty"`
		Quantity int    `json:"quantity"  bson:"quantity"`
		Total    int    `json:"total"  bson:"total"`
	}
)

func (product Product) ToString() string {
	return fmt.Sprintf("id: %s\nName: %s\nPassword: %s\n", product.ProductId, product.Name, product.Price)
}
