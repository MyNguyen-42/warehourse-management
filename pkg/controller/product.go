package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"warehouse-management/pkg/models"
	"warehouse-management/pkg/service"
)

func GetProduct(c echo.Context) error {
	productId := c.Param("id")

	if service.CheckIfProductExists(productId) == false {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Khong tim thay san pham!",
		})
	}

	product, err := service.GetProductById(productId)
	if err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": err,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"product": product,
	})
}

func GetAllProduct(c echo.Context) error {
	products, err := service.GetAllProducts()
	if err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": err,
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"products": products,
	})
}

func CreateProduct(c echo.Context) error {
	param := new(models.Product)
	if err := c.Bind(param); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status": err.Error(),
		})
	}

	productId := param.ProductId

	if service.CheckIfProductExists(productId) == true {
		return c.JSON(http.StatusOK, echo.Map{
			"message":    "Khoi tao san pham that bai, san pham da ton tai!",
			"product_id": productId,
		})
	}

	result, err := service.InsertProduct(param)

	if err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Khoi tao san pham that bai!",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":     "Khoi tao san pham thanh cong",
		"product_id": productId,
		"id":         result.InsertedID,
	})
}

func DeleteProduct(c echo.Context) error {
	productId := c.Param("id")

	if service.CheckIfProductExists(productId) == false {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Khong tim thay san pham!",
		})
	}

	result, err := service.DeleteProduct(productId)
	if err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": err,
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"status": result.DeletedCount,
	})
}

func UpdateProduct(c echo.Context) error {
	param := new(models.Product)
	if err := c.Bind(param); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status": err.Error(),
		})
	}
	productId := c.Param("id")

	if service.CheckIfProductExists(productId) == false {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Khong tim thay san pham!",
		})
	}
	result, err := service.UpdateOneProduct(productId, param)
	if err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": err,
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"result": result,
	})
}
