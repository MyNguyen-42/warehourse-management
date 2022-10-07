package service

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"warehouse-management/pkg/models"
)

const (
	connectionString = "mongodb+srv://mymy:mongodb123@cluster0.xttgkfm.mongodb.net/?retryWrites=true&w=majority"
	dbName           = "warehouse-management"
	colProduct       = "product"
	colWarehouse     = "warehouse"
	colOrders        = "orders"
)

var connection models.Connection

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//config := conf.GetConfig()
	//fmt.Println(config.Database.Hosts[0])
	clientOptions := options.Client().ApplyURI(connectionString)
	//clientOptions := options.Client().ApplyURI("mongodb+srv://" + config.Database.Hosts[0] + "/test?retryWrites=true&w=majority")
	//clientOptions := options.Client().ApplyURI("mongodb://" + config.Database.Hosts[0])

	client, _ := mongo.Connect(ctx, clientOptions)
	connection.ProductCollection = client.Database(dbName).Collection(colProduct)
	connection.WarehouseCollection = client.Database(dbName).Collection(colWarehouse)
	connection.OrdersCollection = client.Database(dbName).Collection(colOrders)
}
