package server

import (
	"github.com/labstack/echo/v4"
	"warehouse-management/pkg/controller"
	"warehouse-management/pkg/framework/api"
	"warehouse-management/pkg/models"
)

func PublicRoutes(group *echo.Group, serverContext *ServiceContext) {
	registerForwardRoutes(group, serverContext.container.Services)
}

func registerForwardRoutes(group *echo.Group, services models.ServicesContainer) {
	api.AddRoute(group, "/warehouse", controller.GetWarehouseNew)
	api.AddRoute(group, "/get-product", services.ProductService.GetProduct)
}
