package services

import (
	"context"
	"warehouse-management/generated/errcode"
	"warehouse-management/interface/services"
	"warehouse-management/pkg/models"
)

func ProductService(container *models.Container) services.ProductService {
	return &productService{
		container,
	}
}

type productService struct {
	*models.Container
}

func (rcv productService) GetProduct(ctx context.Context, req *services.GetProductReq) (*services.GetProductResp, error) {
	routeConfigRepo := rcv.Factory.ProductRepo()
	product, err := routeConfigRepo.FindById(ctx, req.ProductId)

	if err != nil {
		return nil, err
	}
	data := &services.ProductInfo{
		ProductId: product.ProductId,
		Name:      product.Name,
		Price:     product.Price,
	}
	return &services.GetProductResp{
		Product: data,
	}, nil

	return nil, errcode.Error(errcode.RouteNotFound)
}
