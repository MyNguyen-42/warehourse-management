package utils

//
//import (
//	"fmt"
//	"gitlab.ghn.vn/nhanh/backend/work-shift-service/generated/errcode"
//	"go.uber.org/zap"
//)
//
//func ValidateErrorResponse(errResponses map[string]error) (isValid bool) {
//	for _, errResponse := range errResponses {
//		isAppErr := errcode.IsAppError(errResponse)
//		if (isAppErr && errcode.GetAppErrorCode(errResponse) != errcode.OK) || (!isAppErr && errResponse != nil) {
//			return false
//		}
//	}
//	return true
//}
//
//func LoggingError(logger *zap.Logger, action string, err error, errResponseMap map[string]error) {
//	if err != nil {
//		logger.Error(action)
//		return
//	}
//
//	if errResponseMap == nil {
//		return
//	}
//
//	for _, errResponse := range errResponseMap {
//		if errResponse == nil {
//			return
//		}
//
//		if isAppErr := errcode.IsAppError(errResponse); isAppErr && errcode.GetAppErrorCode(errResponse) != errcode.OK || !isAppErr {
//			logger.Error(fmt.Sprintf("%s ", action), zap.Error(errResponse))
//			return
//		}
//	}
//}
