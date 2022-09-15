package models

import "fmt"

type (
	Product struct {
		ProductId string `json:"product_id,omitempty"  bson:"product_id,omitempty"`
		Name      string `json:"name,omitempty"  bson:"name,omitempty"`
		Price     string `json:"price,omitempty"  bson:"price,omitempty"`
	}
)

func (product Product) ToString() string {
	return fmt.Sprintf("id: %s\nName: %s\nPassword: %s\n", product.ProductId, product.Name, product.Price)
}
