package server

//import (
//	"context"
//	"gitlab.ghn.vn/nhanh/backend/work-shift-service/pkg/conf"
//	"gitlab.ghn.vn/nhanh/backend/work-shift-service/pkg/database"
//	"gitlab.ghn.vn/nhanh/backend/work-shift-service/pkg/integration/crm"
//	"gitlab.ghn.vn/nhanh/backend/work-shift-service/pkg/integration/middle_system_api"
//	"gitlab.ghn.vn/nhanh/backend/work-shift-service/pkg/integration/middle_system_api/client"
//	"gitlab.ghn.vn/nhanh/backend/work-shift-service/pkg/integration/nhanh_order"
//	clientNhanhOrder "gitlab.ghn.vn/nhanh/backend/work-shift-service/pkg/integration/nhanh_order/client"
//	"gitlab.ghn.vn/nhanh/backend/work-shift-service/pkg/log"
//	"gitlab.ghn.vn/nhanh/backend/work-shift-service/pkg/models"
//	"gitlab.ghn.vn/nhanh/backend/work-shift-service/pkg/redis"
//	"gitlab.ghn.vn/nhanh/backend/work-shift-service/pkg/repositories"
//	servicesImplement "gitlab.ghn.vn/nhanh/backend/work-shift-service/pkg/services"
//	"go.uber.org/zap"
//)
//
//// NewServiceContext take environment's configurations to initiate a server context
//// which associated database.Database, redis.Client
//func NewServiceContext(appConf *conf.AppConfig) (server *ServiceContext, err error) {
//	ctx := context.Background()
//	var mainDataSource database.Instance
//	var redisClient *redis.Client
//	//var kafkaProducer queue_kafka.KafkaProducer
//	//var workerServer worker.WorkerServer
//
//	// logger associated
//	logger := log.Global()
//	logger.Info("initiating server context...")
//
//	// recover if panic
//	defer func() {
//		if err != nil {
//			logger.Warn("disconnect connections on server context")
//			// disconnect database connection
//			if mainDataSource != nil {
//				mainDataSource.Close()
//			}
//			// close redis connection
//			if redisClient != nil {
//				//_ = redisClient.Close()
//			}
//			//if workerServer != nil {
//			//	workerServer.Shutdown()
//			//}
//		}
//	}()
//
//	logger.Info("setting up main datasource...")
//	// set up databases
//	mainDataSource, err = database.SetupMongoDb(ctx, &appConf.Database.Main)
//	if err != nil {
//		logger.Error("encounter error while setting up main datasource", zap.Error(err))
//		return
//	}
//	logger.Info("set up main datasource successful")
//
//	logger.Info("setting up queue datasource...")
//
//	logger.Info("setting up redis...")
//	if !appConf.RedisSimple {
//		logger.Info("using redis sentinel configuration...")
//		// set up redis client
//		redisClient, err = redis.NewSentinel(ctx, &appConf.RedisSentinelConfig)
//	} else {
//		logger.Info("using redis simple configuration...")
//		// set up redis client
//		redisClient, err = redis.Init(ctx, &appConf.Redis)
//	}
//	if err != nil {
//		logger.Warn("set up redis client failed", zap.Error(err))
//		return
//	} else {
//		logger.Info("set up redis client successful")
//	}
//
//	//workerServer = queue.RabbitMqClient(appConf.QueueConfig, constants.ExchangeQueue)
//	//if err = workerServer.HealthCheck(context.Background()); err != nil {
//	//	logger.Error("failed to health check rabbitmq server", zap.Error(err))
//	//	return
//	//}
//
//	logger.Info("setting up repository factory...")
//	// set up repository factory
//	repositoryFactory := repositories.NewFactory(mainDataSource, appConf.Database.Main.DbName)
//	logger.Info("set up repositories factory successful")
//
//	logger.Info("setting up middle client...")
//	middleApi := client.MiddleSystemClient(appConf.APIPartnerServices.MiddleSystem.Host, appConf.APIPartnerServices.MiddleSystem.AuthToken, appConf.APIPartnerServices.MiddleSystem.Token)
//
//	logger.Info("setting up nhanh order client...")
//	nhanhApi := clientNhanhOrder.NewNhanhClient(appConf.APIPartnerServices.NhanhOrder.Host, appConf.APIPartnerServices.NhanhOrder.AuthToken)
//
//	logger.Info("setting up crm client...")
//	cmrAPI := crm.NewCrmAPI(appConf.APIPartnerServices.MiddleSystem.Host, appConf.APIPartnerServices.MiddleSystem.AuthToken, appConf.APIPartnerServices.MiddleSystem.Token, middleApi)
//
//	//kafkaProducer = queue_kafka.NewKafkaProducer(appConf.KafkaCalculatorQueueConfig.CalculatorProducer)
//
//	return &ServiceContext{
//		mainDataSource: mainDataSource,
//		container: &models.Container{
//			Redis: redisClient,
//			//WorkerServer: workerServer,
//			Factory: repositoryFactory,
//			Api: models.APIContainer{
//				MiddleApi: middleApi,
//				NhanhApi:  nhanhApi,
//				CrmAPI:    cmrAPI,
//			},
//			FileStorageHost:   appConf.APIPartnerServices.MiddleSystem.Host,
//			SystemId:          appConf.SystemId,
//			RegionUnitId:      appConf.RegionUnitId,
//			SecurityJobTitles: appConf.SecurityJobTitles,
//			//KafkaProducer:   &kafkaProducer,
//		},
//		conf: appConf,
//	}, nil
//}
//
//type ServiceContext struct {
//	mainDataSource database.Instance
//	container      *models.Container
//	conf           *conf.AppConfig
//}
//
//// Shutdown close streams
//func (s *ServiceContext) Shutdown() {
//	if s.mainDataSource != nil {
//		s.mainDataSource.Close()
//	}
//	//if s.container.WorkerServer != nil {
//	//	s.container.WorkerServer.Shutdown()
//	//}
//}
//
//func (s *ServiceContext) InitService() {
//	s.container.Services = models.ServicesContainer{
//		MetadataService:         servicesImplement.MetadataService(s.container),
//		IamService:              servicesImplement.IamService(s.conf.APIPartnerServices.Lastmile.Host, s.conf.APIPartnerServices.Lastmile.AuthToken),
//		UserService:             servicesImplement.UserService(s.container),
//		EmployeeService:         servicesImplement.EmployeeService(s.conf.APIPartnerServices.MiddleSystem.Host, s.conf.APIPartnerServices.MiddleSystem.AuthToken),
//		EmployeeCache:           servicesImplement.EmployeeServiceCache(s.container, servicesImplement.EmployeeService(s.conf.APIPartnerServices.MiddleSystem.Host, s.conf.APIPartnerServices.MiddleSystem.AuthToken)),
//		FreeLanceService:        servicesImplement.FreeLanceService(s.conf.APIPartnerServices.Ezjob.Host, s.conf.APIPartnerServices.Ezjob.AuthToken),
//		PermissionService:       servicesImplement.PermissionServer(s.container, s.conf.SystemId, s.conf.Partners, s.conf.Features),
//		ShiftManagementService:  servicesImplement.ShiftManagementService(s.container),
//		UserSessionService:      servicesImplement.UserSessionService(s.container),
//		LayoutManagementService: servicesImplement.LayoutManagementService(s.container),
//		RouteService:            servicesImplement.RouteService(s.container),
//		FileSystemService:       servicesImplement.FileSystemService(s.conf.APIPartnerServices.MiddleSystem.Host, s.conf.APIPartnerServices.MiddleSystem.AuthToken),
//		UploadFileService:       servicesImplement.UploadFileService(s.container),
//		HubManagementService:    servicesImplement.HubManagementService(s.conf.APIPartnerServices.HubService.Host, s.conf.APIPartnerServices.HubService.AuthToken),
//	}
//}
//
//func (s *ServiceContext) Redis() *redis.Client {
//	return s.container.Redis
//}
//
//func (s *ServiceContext) Repositories() repositories.Factory {
//	return s.container.Factory
//}
//
//func (s ServiceContext) MiddleService() middle_system_api.API {
//	return s.container.Api.MiddleApi
//}
//
//func (s ServiceContext) NhanhOrderService() nhanh_order.API {
//	return s.container.Api.NhanhApi
//}
//
////func (s ServiceContext) WorkerServer() worker.WorkerServer {
////	return s.container.WorkerServer
////}
//
//func (s ServiceContext) Container() *models.Container {
//	return s.container
//}
//
//func (s ServiceContext) Config() conf.AppConfig {
//	return *s.conf
//}
