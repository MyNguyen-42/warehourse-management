package service

import (
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"warehouse-management/pkg/models"
)

func getAllWarehouse() ([]models.Warehouse, error) {
	var warehouses []models.Warehouse
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := connection.WarehouseCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var warehouse models.Warehouse
		if err = cursor.Decode(&warehouse); err != nil {
			log.Fatal(err)
		}
		warehouses = append(warehouses, warehouse)
	}

	return warehouses, err
}

func GetWarehouseById(warehouseId string) (models.Warehouse, error) {
	var warehouse models.Warehouse
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{"warehouse_id": warehouseId}
	err := connection.WarehouseCollection.FindOne(ctx, filter).Decode(&warehouse)
	products := warehouse.Products
	for index, element := range products {
		product, _ := GetProductById(element.ProductId)
		products[index] = *product
		products[index].Quantity = element.Quantity
	}
	return warehouse, err
}

func updateQuantityOneProductInWarehouse(warehouseId string, productId string, quantity int) {
	filter := bson.M{"warehouse_id": warehouseId}
	update := bson.M{"$set": bson.M{"products.$[elem].quantity": quantity}}
	connection.WarehouseCollection.FindOneAndUpdate(
		context.Background(),
		filter,
		update,
		options.FindOneAndUpdate().SetArrayFilters(options.ArrayFilters{
			Filters: []interface{}{bson.M{"elem.product_id": productId}},
		}),
	)
}

func addProductToWarehouse(warehouseId string, product *models.Product, quantity int) bool {
	filter := bson.M{"warehouse_id": warehouseId}
	if !CheckIfProductExists(product.ProductId) {
		_, err := InsertProduct(product)
		if err != nil {
			return false
		}
	}
	update := bson.M{"$push": bson.M{"products": bson.M{"product_id": product.ProductId, "quantity": quantity}}}
	connection.WarehouseCollection.UpdateOne(context.Background(), filter, update)
	return true
}

func isExitOnListProduct(product *models.Product, products []models.Product) (isExit bool, indexOfItem int) {
	for index, element := range products {
		if element.ProductId == product.ProductId {
			return true, index
		}
	}
	return false, -1
}

func InsertOrder(order *models.Warehouse) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	inserted, err := connection.OrdersCollection.InsertOne(ctx, order)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in db with id", inserted.InsertedID)
	return inserted, err
}

func ImportProductIntoWarehouse(warehouseId string, products []models.Product) bool {
	warehouse, _ := GetWarehouseById(warehouseId)
	for _, element := range products {
		isExit, index := isExitOnListProduct(&element, warehouse.Products)
		if isExit {
			quantity := warehouse.Products[index].Quantity + element.Quantity
			warehouse.Products[index].Quantity = quantity
			updateQuantityOneProductInWarehouse(warehouseId, warehouse.Products[index].ProductId, quantity)
		} else {
			warehouse.Products = append(warehouse.Products, element)
			addProductToWarehouse(warehouseId, &element, element.Quantity)
		}
	}
	fmt.Println("warehouse before imported: ", warehouse)
	return true
}

func ExportProductFromWarehouse(warehouseId string, products []models.Product) (bool, []models.Product) {
	warehouse, _ := GetWarehouseById(warehouseId)
	for _, element := range products {
		isExit, index := isExitOnListProduct(&element, warehouse.Products)
		if isExit {
			quantity := warehouse.Products[index].Quantity - element.Quantity
			if quantity < 0 {
				return false, products
			}
		} else {
			return false, products
		}
	}
	for _, element := range products {
		isExit, index := isExitOnListProduct(&element, warehouse.Products)
		if isExit {
			quantity := warehouse.Products[index].Quantity - element.Quantity
			warehouse.Products[index].Quantity = quantity
			total := element.Quantity * element.Price
			warehouse.Products[index].Total = total
			updateQuantityOneProductInWarehouse(warehouseId, warehouse.Products[index].ProductId, quantity)
		}
	}
	fmt.Println("warehouse before imported: ", warehouse)
	return true, products
}

//func ExportProductFromWarehouseForOrder(warehouseId string, products []models.Product) (bool, []models.OrderProductParams) {
//	warehouse, _ := GetWarehouseById(warehouseId)
//	productToReturn:= []models.OrderProductParams
//	for _, element := range products {
//		isExit, index := isExitOnListProduct(&element, warehouse.Products)
//		if isExit {
//			quantity := warehouse.Products[index].Quantity - element.Quantity
//			if quantity < 0 {
//				productToReturn[index].Name= warehouse.Products[index].Name
//				return false, productToReturn
//			}
//		} else {
//			return false, productToReturn
//		}
//	}
//	for _, element := range products {
//		isExit, index := isExitOnListProduct(&element, warehouse.Products)
//		if isExit {
//			quantity := warehouse.Products[index].Quantity - element.Quantity
//			warehouse.Products[index].Quantity = quantity
//			total := element.Quantity * element.Price
//			warehouse.Products[index].Total = total
//			updateQuantityOneProductInWarehouse(warehouseId, warehouse.Products[index].ProductId, quantity)
//		}
//	}
//	fmt.Println("warehouse before imported: ", warehouse)
//	return true, products
//}
func CheckIfWarehouseExists(warehouseId string) bool {
	products, _ := getAllWarehouse()
	for _, element := range products {
		if element.WarehouseId == warehouseId {
			return true
		}
	}
	return false
}
