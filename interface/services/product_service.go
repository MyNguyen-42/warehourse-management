package services

import (
	"context"
)

type ProductService interface {
	GetProduct(context.Context, *GetProductReq) (*GetProductResp, error)
	AddProduct(context.Context, *AddProductReq) (*AddProductResp, error)
	EditProduct(context.Context, *EditProductReq) (*EditProductResp, error)
}

type GetProductReq struct {
	ProductId string `json:"product_id"`
}

type GetProductResp struct {
	Product *ProductInfo `json:"product"`
}
type AddProductReq struct {
}

type AddProductResp struct {
}

type EditProductReq struct {
}

type EditProductResp struct {
}

type ProductInfo struct {
	ProductId string `json:"product_id"`
	Name      string `json:"name"`
	Price     int    `json:"price"`
}
