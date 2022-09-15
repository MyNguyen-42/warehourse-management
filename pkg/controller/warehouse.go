package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"warehouse-management/pkg/models"
	"warehouse-management/pkg/service"
)

func GetWarehouse(c echo.Context) error {
	warehouseId := c.Param("id")
	if service.CheckIfWarehouseExists(warehouseId) == false {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Kho khong ton tai!",
		})
	}
	warehouse, _ := service.GetWarehouseById(warehouseId)

	return c.JSON(http.StatusOK, echo.Map{
		"warehouse": warehouse,
	})
}

func ImportIntoWarehouse(c echo.Context) error {
	param := new(models.Warehouse)
	if err := c.Bind(param); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status": err.Error(),
		})
	}

	if service.CheckIfWarehouseExists(param.WarehouseId) == false {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Kho khong ton tai!",
		})
	}

	isSuccess := service.ImportProductIntoWarehouse(param.WarehouseId, param.Products)
	if isSuccess {
		return c.JSON(http.StatusOK, echo.Map{"message": "import success!"})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "import fail!"})
}

func RotationProductWarehouse(c echo.Context) error {
	param := new(models.RotaWarehouseRequest)
	if err := c.Bind(param); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status": err.Error(),
		})
	}

	if service.CheckIfWarehouseExists(param.FromWarehouseId) == false || service.CheckIfWarehouseExists(param.ToWarehouseId) == false {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Kho " + param.FromWarehouseId + "hoac" + param.ToWarehouseId + "khong ton tai!",
		})
	}

	if success, _ := service.ExportProductFromWarehouse(param.FromWarehouseId, param.Products); success {
		service.ImportProductIntoWarehouse(param.ToWarehouseId, param.Products)
		return c.JSON(http.StatusOK, echo.Map{"status": "rotation success!"})
	}
	return c.JSON(http.StatusOK, echo.Map{"status": "rotation fail!"})
}

type OrderRequest struct {
	models.Warehouse
	Total int `json:"total,omitempty"`
}

func Order(c echo.Context) error {
	param := new(models.Warehouse)
	if err := c.Bind(param); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status": err.Error(),
		})
	}
	isSuccess, products := service.ExportProductFromWarehouse(param.WarehouseId, param.Products)
	param.Products = products
	service.InsertOrder(param)
	if isSuccess {
		return c.JSON(http.StatusOK, echo.Map{"message": "order success!"})
		fmt.Println(products)
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "order fail! warehouse cannot have enough quantity of product"})

}
