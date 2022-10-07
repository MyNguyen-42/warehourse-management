package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

const connectionString = "mongodb+srv://mymy:mongodb123@cluster0.xttgkfm.mongodb.net/?retryWrites=true&w=majority"

func SetupMongoDb(ctx context.Context) (Instance, error) {
	//logger := log.Logger(ctx)
	//logger.Sugar().Infof("initializing connection to mongodb, username: %s, uri: %s, replSet: %s, hosts: %s, authSource: %s",
	//	config.Username, strings.TrimSpace(config.URI), config.ReplSet, strings.Join(config.Hosts, ","), config.AuthSource)

	//clientOpts := options.Client().
	//	SetHosts(config.Hosts).
	//	SetAuth(options.Credential{
	//		AuthSource:  config.AuthSource,
	//		Username:    config.Username,
	//		Password:    config.Password,
	//		PasswordSet: true,
	//	})

	clientOpts := options.Client().ApplyURI(connectionString)

	//if config.ReplSet != "" {
	//	clientOpts = clientOpts.SetReplicaSet(config.ReplSet)
	//}
	//if strings.TrimSpace(config.URI) != "" {
	//	clientOpts = clientOpts.ApplyURI(config.URI)
	//}

	// set timeout to connect mongodb after 10 seconds
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	fmt.Println("connecting to mongodb...")

	// connect
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, err
	}
	fmt.Println("connected to mongodb")

	fmt.Println("ping mongodb...")
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("ping mongodb err...")
		return nil, err
	}
	fmt.Printf("ping success mongodb host: %s", clientOpts.Hosts)

	return newInstance(client, true), nil
}

type Config struct {
	URI              string   `json:"uri" mapstructure:"uri"`
	Hosts            []string `json:"hosts" mapstructure:"hosts"`
	AuthSource       string   `json:"authSource" mapstructure:"authSource"`
	Username         string   `json:"username" mapstructure:"username"`
	Password         string   `json:"password" mapstructure:"password"`
	ReplSet          string   `json:"replSet" mapstructure:"replSet"`
	DbName           string   `json:"dbName" mapstructure:"dbName"`
	UnUseTransaction bool     `json:"unUseTransaction" mapstructure:"unUseTransaction"`
}
