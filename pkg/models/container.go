package models

import (
	"warehouse-management/interface/services"
	"warehouse-management/pkg/repositories"
)

type Container struct {
	Factory           repositories.Factory
	Api               APIContainer
	Services          ServicesContainer
	FileStorageHost   string
	SystemId          int
	RegionUnitId      int
	SecurityJobTitles []int
}

type APIContainer struct {
	//MiddleApi middle_system_api.API
	//NhanhApi  nhanh_order.API
	//CrmAPI    crm.CrmAPI
}

type ServicesContainer struct {
	ProductService services.ProductService
}
