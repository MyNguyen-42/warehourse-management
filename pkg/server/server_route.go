package server

//
//import (
//	"gitlab.ghn.vn/nhanh/backend/work-shift-service/pkg/framework/api"
//	"gitlab.ghn.vn/nhanh/backend/work-shift-service/pkg/models"
//	"net/http"
//	"time"
//
//	"github.com/labstack/echo/v4"
//	"gitlab.ghn.vn/nhanh/backend/work-shift-service/pkg/constants"
//	"gitlab.ghn.vn/nhanh/backend/work-shift-service/pkg/middlewares"
//	"gitlab.ghn.vn/nhanh/backend/work-shift-service/pkg/models/common"
//)
//
//func PublicRoutes(group *echo.Group, serverContext *ServiceContext) {
//	registerMetadataRoutes(group, serverContext.container.Services)
//	registerCheckInRoutes(group, serverContext.container.Services)
//	registerLayoutRoutes(group, serverContext.container.Services)
//	registerScheduleRoutes(group, serverContext.container.Services)
//	registerSessionRoutes(group, serverContext.container.Services)
//	registerRouteRoutes(group, serverContext.container.Services)
//	registerForwardRoutes(group, serverContext.container.Services)
//	registerDashboardRoutes(group, serverContext.container.Services)
//	registerImportRoutes(group, serverContext.container.Services)
//	registerDebugRoutes(group, serverContext)
//}
//
//func PrivateRoutes(group *echo.Group, serverContext *ServiceContext) {
//	registerMetadataRoutes(group, serverContext.container.Services)
//	registerCheckInRoutes(group, serverContext.container.Services)
//	registerSessionRoutes(group, serverContext.container.Services)
//	registerRouteRoutes(group, serverContext.container.Services)
//	registerForwardRoutes(group, serverContext.container.Services)
//}
//
//func registerMetadataRoutes(group *echo.Group, services models.ServicesContainer) {
//	api.AddRoute(group, "/metadata/get-user-types", services.MetadataService.GetUserTypes)
//	api.AddRoute(group, "/metadata/get-hub-actions", services.MetadataService.GetHubActions)
//	api.AddRoute(group, "/metadata/get-checkin-type", services.MetadataService.GetCheckinTypes)
//	api.AddRoute(group, "/metadata/get-hub-area-action", services.MetadataService.GetHubAreaAction)
//	api.AddRoute(group, "/metadata/get-route-conditions", services.MetadataService.GetRouteConditions)
//	api.AddRoute(group, "/metadata/all-warehouses", services.HubManagementService.GetAllHubs)
//	api.AddRoute(group, "/metadata/gen-ezjob-download-token", services.MetadataService.GenEzjobDownloadToken)
//}
//
//func registerRouteRoutes(group *echo.Group, services models.ServicesContainer) {
//	api.AddRoute(group, "/route/distribute-sound", services.RouteService.FindRoute)
//	api.AddRoute(group, "/route/get-list-route", services.RouteService.GetListRoute)
//	api.AddRoute(group, "/route/import", services.RouteService.ImportRouteConfig)
//	api.AddRoute(group, "/route/gen-upload-token", services.UploadFileService.GenerateTokenUploadFile)
//	api.AddRoute(group, "/route/gen-download-token", services.UploadFileService.GenerateTokenDownloadFile)
//	api.AddRoute(group, "/route/update-upload-token", services.UploadFileService.GenerateTokenUploadFile)
//}
//
//func registerForwardRoutes(group *echo.Group, services models.ServicesContainer) {
//	api.AddRoute(group, "/forward/import-version", services.RouteService.ImportForwardVersion)
//	api.AddRoute(group, "/forward/apply-version", services.RouteService.ApplyForwardVersion)
//	api.AddRoute(group, "/forward/delete-version", services.RouteService.DeleteForwardVersion)
//	api.AddRoute(group, "/forward/get-versions", services.RouteService.GetForwardVersions)
//	api.AddRoute(group, "/forward/get-version", services.RouteService.GetForwardVersion)
//	api.AddRoute(group, "/forward/get-forward-hub", services.RouteService.GetForwardHub)
//}
//
//func registerCheckInRoutes(group *echo.Group, services models.ServicesContainer) {
//	api.AddRoute(group, "/generate-my-qr", services.UserSessionService.GenerateMyQr)
//	api.AddRoute(group, "/check-hub-member", services.UserSessionService.CheckHubMember)
//	api.AddRoute(group, "/check-in", services.UserSessionService.SelfCheckin)
//	api.AddRoute(group, "/scan-member", services.UserSessionService.ScanMember)
//	api.AddRoute(group, "/my-collaborators", services.UserSessionService.MyCollaborators)
//	api.AddRoute(group, "/current-session", services.UserSessionService.CurrentSession)
//	api.AddRoute(group, "/checkin-history", services.UserSessionService.UserCheckinHistory)
//}
//
//func registerSessionRoutes(group *echo.Group, services models.ServicesContainer) {
//	api.AddRoute(group, "/session/get-collaborators", services.UserSessionService.GetCollaborators)
//	api.AddRoute(group, "/session/gen-session", services.UserSessionService.GenerateSession)
//	api.AddRoute(group, "/session/generate-qr", services.UserSessionService.GenerateQr)
//	api.AddRoute(group, "/session/get-session", services.UserSessionService.GetSession)
//	api.AddRoute(group, "/session/switch-area", services.UserSessionService.SwitchAreaUserSession)
//	api.AddRoute(group, "/session/force-check-out", services.UserSessionService.Checkout)
//}
//
//func registerDashboardRoutes(group *echo.Group, services models.ServicesContainer) {
//	api.AddRoute(group, "/dashboard/get-users-in-area", services.UserSessionService.GetUsersInArea)
//	api.AddRoute(group, "/dashboard/list-area-checkin", services.UserSessionService.ListAreaCheckin)
//	api.AddRoute(group, "/dashboard/checkin-history", services.UserSessionService.ListUserCheckinHistory)
//}
//
//func registerLayoutRoutes(group *echo.Group, services models.ServicesContainer) {
//	api.AddRoute(group, "/layout/create-layout", services.LayoutManagementService.CreateLayout)
//	api.AddRoute(group, "/layout/apply-layout", services.LayoutManagementService.ApplyLayout)
//	api.AddRoute(group, "/layout/delete-layout", services.LayoutManagementService.DeleteLayout)
//	api.AddRoute(group, "/layout/get-layout-list", services.LayoutManagementService.GetLayoutList)
//	api.AddRoute(group, "/layout/create-area", services.LayoutManagementService.CreateArea)
//	api.AddRoute(group, "/layout/get-area-list", services.LayoutManagementService.GetAreaList)
//	api.AddRoute(group, "/layout/update-area", services.LayoutManagementService.UpdateArea)
//	api.AddRoute(group, "/layout/active-area", services.LayoutManagementService.SwitchActive)
//	api.AddRoute(group, "/layout/get-area", services.LayoutManagementService.GetArea)
//	api.AddRoute(group, "/layout/get-hub-config", services.LayoutManagementService.GetHubConfig)
//	api.AddRoute(group, "/layout/get-area-history", services.LayoutManagementService.GetAreaHistory)
//	api.AddRoute(group, "/layout/refresh-checkin-code", services.LayoutManagementService.RefreshCheckinCode)
//	api.AddRoute(group, "/layout/delete-area", services.LayoutManagementService.DeleteArea)
//}
//
//func registerScheduleRoutes(group *echo.Group, services models.ServicesContainer) {
//	api.AddRoute(group, "/shift/create-shift", services.ShiftManagementService.CreateShift)
//	api.AddRoute(group, "/shift/get-shifts", services.ShiftManagementService.GetShifts)
//	api.AddRoute(group, "/shift/update-shifts", services.ShiftManagementService.UpdateShifts)
//	api.AddRoute(group, "/shift/delete-shift", services.ShiftManagementService.DeleteShift)
//	api.AddRoute(group, "/shift/delete-schedule", services.ShiftManagementService.DeleteSchedule)
//	api.AddRoute(group, "/shift/get-hub-users", services.UserService.GetHubUsers)
//}
//
//func registerImportRoutes(group *echo.Group, services models.ServicesContainer) {
//	group.Add(http.MethodPost, "/shift/import", services.ShiftManagementService.ImportExcelFile)
//}
//
//func registerDebugRoutes(group *echo.Group, serverContext *ServiceContext) {
//	debugRouter := group.Group("/debug")
//	if serverContext.conf.Env == constants.EnvLocal {
//		debugRouter.Use(addUserActionLocal)
//	}
//}
//
//func addUserActionLocal(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
//	return func(c echo.Context) error {
//		req := c.Request()
//		ctx := c.Request().Context()
//
//		ctx = middlewares.WithUserAction(ctx, &common.ActionSource{
//			Id:     constants.SystemUser,
//			Source: "debugger",
//			Time:   time.Now(),
//		})
//
//		c.SetRequest(req.WithContext(ctx))
//		return handlerFunc(c)
//	}
//}
