package server

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"warehouse-management/pkg/conf"
	"warehouse-management/pkg/database"
	"warehouse-management/pkg/log"
	"warehouse-management/pkg/models"
	"warehouse-management/pkg/repositories"
	"warehouse-management/pkg/services"
)

// NewServiceContext take environment's configurations to initiate a server context
// which associated database.Database, redis.Client
func NewServiceContext() (server *ServiceContext, err error) {
	ctx := context.Background()
	var mainDataSource database.Instance

	// logger associated
	logger := log.Global()
	fmt.Println("initiating server context...")

	// recover if panic
	defer func() {
		if err != nil {
			logger.Warn("disconnect connections on server context")
			if mainDataSource != nil {
				mainDataSource.Close()
			}
		}
	}()

	fmt.Println("setting up main datasource...")
	// set up databases
	mainDataSource, err = database.SetupMongoDb(ctx)
	if err != nil {
		logger.Error("encounter error while setting up main datasource", zap.Error(err))
		return
	}
	fmt.Println("set up main datasource successful")

	fmt.Println("setting up queue datasource...")

	fmt.Println("setting up redis...")
	if err != nil {
		logger.Warn("set up redis client failed", zap.Error(err))
		return
	} else {
		fmt.Println("set up redis client successful")
	}

	fmt.Println("setting up repository factory...")
	repositoryFactory := repositories.NewFactory(mainDataSource, "warehouse-management")
	fmt.Println("set up repositories factory successful")

	return &ServiceContext{
		mainDataSource: mainDataSource,
		container: &models.Container{
			Factory: repositoryFactory,
			Api:     models.APIContainer{},
		},
	}, nil
}

type ServiceContext struct {
	mainDataSource database.Instance
	container      *models.Container
	conf           *conf.Config
}

// Shutdown close streams
func (s *ServiceContext) Shutdown() {
	if s.mainDataSource != nil {
		s.mainDataSource.Close()
	}
	//if s.container.WorkerServer != nil {
	//	s.container.WorkerServer.Shutdown()
	//}
}

func (s *ServiceContext) InitService() {
	s.container.Services = models.ServicesContainer{
		ProductService: services.ProductService(s.container),
	}
}

func (s *ServiceContext) Repositories() repositories.Factory {
	return s.container.Factory
}

func (s ServiceContext) Container() *models.Container {
	return s.container
}

func (s ServiceContext) Config() conf.Config {
	return *s.conf
}
