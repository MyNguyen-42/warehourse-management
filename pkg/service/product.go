package service

import (
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"warehouse-management/pkg/models"
)

func InsertProduct(product *models.Product) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	productToInsert := models.InsertProductParams{
		ProductId: product.ProductId,
		Name:      product.Name,
		Price:     product.Price,
	}

	inserted, err := connection.ProductCollection.InsertOne(ctx, productToInsert)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in db with id", inserted.InsertedID)
	return inserted, err
}

func UpdateOneProduct(productId string, product *models.Product) (*mongo.UpdateResult, error) {
	filter := bson.M{"product_id": productId}
	productToUpdate := models.UpdateProductParams{
		Name:  product.Name,
		Price: product.Price,
	}
	update := bson.D{{"$set", productToUpdate}}
	result, err := connection.ProductCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count: ", result.ModifiedCount)
	return result, err
}

func DeleteProduct(productId string) (*mongo.DeleteResult, error) {
	filter := bson.M{"product_id": productId}
	deleteCount, err := connection.ProductCollection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("delete count: ", deleteCount.DeletedCount)
	return deleteCount, err
}

func GetAllProducts() ([]models.Product, error) {
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

	return products, err
}

func GetProductById(productId string) (*models.Product, error) {
	var product models.Product
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{"product_id": productId}
	err := connection.ProductCollection.FindOne(ctx, filter).Decode(&product)
	return &product, err
}

func CheckIfProductExists(productId string) bool {
	products, _ := GetAllProducts()
	for _, element := range products {
		if element.ProductId == productId {
			return true
		}
	}
	return false
}
