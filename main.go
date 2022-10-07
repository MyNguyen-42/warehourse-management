package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"warehouse-management/pkg/conf"
	"warehouse-management/pkg/server"
	"warehouse-management/pkg/server/protocols/http"
)

func init() {
	conf.LoadConfigFiles(".")
}

func main() {
	config := conf.GetConfig()

	gServer, err := server.NewServiceContext()

	if err != nil {
		fmt.Println("initiate server context failed", zap.Error(err))
	}
	fmt.Println(gServer)
	gServer.InitService()

	httpServer := http.EchoServer()
	serviceGroup := httpServer.Group("/v1")
	AddRoutes(serviceGroup, gServer)

	go func() {
		if err := httpServer.Start(fmt.Sprintf(":%d", config.Port)); err != nil {
			fmt.Println("server shutdown")
		}
	}()

	// graceful shutdown
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGTERM, syscall.SIGINT)
	<-shutdown

	//e := echo.New()
	//
	//e.GET("/", func(c echo.Context) error {
	//	return c.String(http.StatusOK, "Hello, World!")
	//})
	//e.GET("/products", controller.GetAllProduct)
	//e.GET("/product/:id", controller.GetProduct)
	//e.PUT("/product/:id", controller.UpdateProduct)
	//e.POST("/product", controller.CreateProduct)
	//e.DELETE("product/:id", controller.DeleteProduct)
	//
	//e.GET("/warehouse/:id", controller.GetWarehouse)
	//
	//e.POST("/import-in-warehouse", controller.ImportIntoWarehouse)
	//
	//e.POST("/rotation-warehouse", controller.RotationProductWarehouse)
	//
	//e.POST("/order", controller.Order)
	//
	//e.Logger.Fatal(e.Start(":8080"))
}

func AddRoutes(serviceGroup *echo.Group, gServer *server.ServiceContext) {
	publicGroup := serviceGroup.Group("/public")
	server.PublicRoutes(publicGroup, gServer)

}
