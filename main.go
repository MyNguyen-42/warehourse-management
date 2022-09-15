package main

import (
	"github.com/labstack/echo/v4"
	//"warehouse-management/pkg/server/protocols/http"

	"net/http"
	"warehouse-management/pkg/controller"
)

func main() {
	//httpServer := http.EchoServer(&appConf)
	//serviceGroup := httpServer.Group(constants.ServiceSubDomain)
	//AddRoutes(serviceGroup, gServer)

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/products", controller.GetAllProduct)
	e.GET("/product/:id", controller.GetProduct)
	e.PUT("/product/:id", controller.UpdateProduct)
	e.POST("/product", controller.CreateProduct)
	e.DELETE("product/:id", controller.DeleteProduct)

	e.GET("/warehouse/:id", controller.GetWarehouse)

	e.POST("/import-in-warehouse", controller.ImportIntoWarehouse)

	e.POST("/rotation-warehouse", controller.RotationProductWarehouse)

	e.POST("/order", controller.Order)

	e.Logger.Fatal(e.Start(":8080"))
}

//func AddRoutes(serviceGroup *echo.Group, gServer *server.ServiceContext) {
//	// Define group
//	publicGroup := serviceGroup.Group("/public")
//	server.PublicRoutes(publicGroup, gServer)
//
//	internalGroup := serviceGroup.Group("/internal")
//	server.PrivateRoutes(internalGroup, gServer)
//}
