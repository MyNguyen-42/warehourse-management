package controller

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"time"
	"warehouse-management/models"
)

const connectionString = "mongodb+srv://mymy:mongodb123@cluster0.xttgkfm.mongodb.net/?retryWrites=true&w=majority"
const dbName = "warehouse-management"
const colName = "product"

var connection models.Connection

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(connectionString)
	client, _ := mongo.Connect(ctx, clientOptions)
	connection.ProductCollection = client.Database(dbName).Collection(colName)
}

func insertProduct(product models.Product) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	inserted, err := connection.ProductCollection.InsertOne(ctx, product)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in db with id", inserted.InsertedID)
}

func updateOneProduct(productId string, name string) {
	filter := bson.D{{"product_id", productId}}
	update := bson.D{{"$set", bson.D{{"name", name}}}}
	result, err := connection.ProductCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count: ", result.ModifiedCount)
}

func deleteOneProduct(productId string) {
	filter := bson.M{"product_id": productId}
	deleteCount, err := connection.ProductCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("delete count: ", deleteCount.DeletedCount)
}

func deleteAllProduct() {
	deleteResult, err := connection.ProductCollection.DeleteOne(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Number of movies delete: ", deleteResult.DeletedCount)
}

func getAllProducts() []models.Product {
	var products []models.Product
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := connection.ProductCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var product models.Product
		if err = cursor.Decode(&product); err != nil {
			log.Fatal(err)
		}
		products = append(products, product)
	}

	return products
}

func GetAllProduct(c echo.Context) error {
	products := getAllProducts()
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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	products := getAllProducts()

	for i := 0; i < len(products); i++ {
		fmt.Println(products[i].ProductId, param.ProductId)
		if products[i].ProductId == param.ProductId {
			return c.JSON(http.StatusOK, echo.Map{
				"status":     "Khoi tao san pham that bai, san pham da ton tai",
				"product_id": param.ProductId,
			})
		}
	}

	result, err := connection.ProductCollection.InsertOne(ctx, param)
	if err != nil {
		panic(err)
	}

	id := result.InsertedID

	return c.JSON(http.StatusOK, echo.Map{
		"status":     "Khoi tao san pham thanh cong",
		"product_id": param.ProductId,
		"id":         id,
	})
}

func DeleteProduct(c echo.Context) error {
	productId := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := connection.ProductCollection.DeleteMany(ctx, bson.M{"product_id": productId})
	if err != nil {
		log.Fatal(err)
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

	filter := bson.D{{"product_id", productId}}
	update := bson.D{{"$set", bson.D{{"name", param.Name}}}}
	result, err := connection.ProductCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, echo.Map{
		"result": result,
	})
}

func GetProduct(c echo.Context) error {
	productId := c.Param("id")
	var products []models.Product
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := connection.ProductCollection.Find(ctx, bson.D{{"product_id", productId}})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var product models.Product
		if err = cursor.Decode(&product); err != nil {
			log.Fatal(err)
		}
		products = append(products, product)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"products": products,
	})
}
