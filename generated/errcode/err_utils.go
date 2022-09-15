package errcode

import (
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"strings"
)

const errorPrefix = "#"
const errorSuffix = " - "

func Error(c ErrorCode, args ...interface{}) error {
	return ErrorWithData(c, nil, args...)
}

func ErrorWithData(c ErrorCode, data interface{}, args ...interface{}) error {
	message := c.String()
	if strings.Count(message, "%") > 0 {
		message = fmt.Sprintf(c.String(), args...)
	}
	return &AppError{
		Code:    c,
		Data:    data,
		Message: message,
	}
}

func ErrorWithMessage(c ErrorCode, data interface{}, message string) error {
	return &AppError{
		Code:    c,
		Data:    data,
		Message: message,
	}
}

type AppError struct {
	Code    ErrorCode   `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (err *AppError) Error() string {
	return fmt.Sprintf("%s%s", codeFormat(err.Code), err.Message)
}

//IsAppError using reflection to verify whether the objectPtr is app error, returns false if objectPtr is nil
func IsAppError(objectPtr interface{}) bool {
	if objectPtr == nil {
		return false
	}

	var ptr reflect.Value
	var elem reflect.Value
	elem = reflect.ValueOf(objectPtr)
	if elem.Type().Kind() == reflect.Ptr {
		ptr = elem
		elem = ptr.Elem() // acquire value referenced by pointer
	} else {
		ptr = reflect.New(reflect.TypeOf(objectPtr)) // create new pointer
		temp := ptr.Elem()                           // create variable to value of pointer
		temp.Set(elem)                               // set value of variable to our passed in value
	}
	return elem.Type() == reflect.TypeOf(AppError{})
}

func codeFormat(code ErrorCode) string {
	return fmt.Sprintf("%s%d%s", errorPrefix, code, errorSuffix)
}

func GetAppErrorCode(err interface{}) ErrorCode {
	if IsAppError(err) {
		status := err.(*AppError).Code
		return status
	}
	return Unknown
}

func GetAppErrorMessage(err interface{}) string {
	if IsAppError(err) {
		status := err.(*AppError).Message
		if status != "" {
			return status
		}
	}
	return Unknown.Error()
}

func GetAppErrorData(err interface{}) interface{} {
	if IsAppError(err) {
		return err.(*AppError).Data
	}
	return nil
}

func IsNotFoundError(err error) bool {
	return errors.Is(err, mongo.ErrNoDocuments) || GetAppErrorCode(err) == NotFound
}
