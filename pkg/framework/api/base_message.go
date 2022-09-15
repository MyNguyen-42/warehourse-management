package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"warehouse-management/generated/errcode"
)

type BaseMessage struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ResponseOK API response error code OK and Message of OK with data was attached
func ResponseOK(c echo.Context, data interface{}) error {
	resp := &BaseMessage{
		Code:    int(errcode.OK),
		Message: errcode.OK.Error(),
		Data:    data,
	}
	return c.JSON(http.StatusOK, resp)
}

// ResponseError response with internal error code
func ResponseError(c echo.Context, err error) error {
	if errcode.IsNotFoundError(err) {
		err = errcode.Error(errcode.NotFound)
	}
	resp := &BaseMessage{
		Data: nil,
	}
	httpStatus := http.StatusBadRequest
	// thay thế lỗi lib echo bằng lỗi app
	if he, ok := err.(*echo.HTTPError); ok {
		if he.Code == http.StatusNotFound {
			err = errcode.Error(errcode.NotFound)
			httpStatus = http.StatusNotFound
		}
	}
	resp.Code = int(errcode.GetAppErrorCode(err))
	resp.Data = errcode.GetAppErrorData(err)
	resp.Message = errcode.GetAppErrorMessage(err)
	return c.JSON(httpStatus, resp)
}
