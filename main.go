package main

import (
	"github.com/labstack/echo/v4"
	"warehouse-management/controller"
)

func main() {
	e := echo.New()

	e.GET("/products", controller.GetAllProduct)
	e.GET("/product/:id", controller.GetProduct)
	e.POST("/product", controller.CreateProduct)
	e.PUT("/product/:id", controller.UpdateProduct)
	e.DELETE("product/:id", controller.DeleteProduct)

	e.Logger.Fatal(e.Start(":8080"))
}
