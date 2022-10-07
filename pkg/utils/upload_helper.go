package utils

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"
	"warehouse-management/generated/errcode"
	"warehouse-management/pkg/log"
)

type FileStorageData struct {
	FileId        string
	FileName      string
	FileExtension string
	Path          string
	Data          []byte
}

func UploadFile(c echo.Context, uploadType string) (*FileStorageData, error) {
	fileId := primitive.NewObjectID().Hex()
	fileName, fileData, err := GetAttachments(c, "attachments")
	if err != nil {
		return nil, err
	}

	filePath := joinPath("Upload", uploadType, fileId)

	dst, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	if _, err = dst.Write(fileData); err != nil {
		return nil, err
	}

	return &FileStorageData{
		FileId:        fileId,
		FileName:      fileName,
		Path:          filePath,
		Data:          fileData,
		FileExtension: filepath.Ext(fileName),
	}, nil
}

func GetAttachments(ctx echo.Context, key string) (string, []byte, error) {
	var request = ctx.Request()
	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	err := request.ParseMultipartForm(10 << 20)
	if err != nil {
		log.Logger(ctx.Request().Context()).Error("Parse error: " + err.Error())
		return "", nil, err
	}

	file, header, err := request.FormFile(key)
	if err != nil {
		log.Logger(ctx.Request().Context()).Error("read file error: " + err.Error())
		return "", nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Logger(ctx.Request().Context()).Error("read file error: " + err.Error())
		return "", nil, err
	}
	return header.Filename, data, nil
}

func joinPath(paths ...string) (filePath string) {
	for _, path := range paths {
		filePath = filepath.Join(filePath, path)
		if _, err := os.Stat(filepath.Dir(filePath)); os.IsNotExist(err) {
			_ = os.Mkdir(filepath.Dir(filePath), 0775)
		}
	}
	return
}

func MapObjectByTagNew(ptr interface{}, dict map[string]string, tag string) error {
	var structValue = reflect.ValueOf(ptr)

	// validate struct value
	if structValue.Kind() != reflect.Ptr || structValue.IsNil() {
		return errcode.Error(errcode.InvalidImportFile, "can not map to nil ptr")
	} else {
		structValue = structValue.Elem()
	}

	// validate struct type
	var structType = reflect.TypeOf(ptr).Elem()
	if structType.Kind() != reflect.Struct {
		return errcode.Error(errcode.InvalidImportFile, "can not map to non struct type")
	}

	var numField = structType.NumField()
	for fi := 0; fi < numField; fi++ {
		var fieldType = structType.Field(fi)
		var fieldValue = structValue.Field(fi)

		var key = fieldType.Tag.Get(tag)
		var val, ok = dict[key]
		if strings.TrimSpace(key) == "" || !ok {
			continue
		}

		if !fieldValue.CanSet() {
			return fmt.Errorf("field %s unsettable", key)
		}

		_ = setProperlyValue(fieldValue.Type().Kind(), fieldValue, val)
	}
	return nil
}

func setProperlyValue(k reflect.Kind, v reflect.Value, data string) (err error) {
	if v.Type() == reflect.TypeOf(time.Now()) {
		if input, err := ParseTime(data); err != nil {
			return err
		} else {
			v.Set(reflect.ValueOf(input))
		}
	}
	switch k {
	case reflect.Int, reflect.Int64:
		var intData int64
		var dot = strings.Index(data, ".")
		var intStr = data
		if dot >= 0 {
			intStr = data[:dot]
			if len(data) > len(intStr)+1 {
				var precision = data[dot+1:]
				if precisionData, parseErr := strconv.ParseInt(precision, 10, 64); parseErr != nil || precisionData > 0 {
					err = fmt.Errorf("can't set %s to type %s", data, v.Type().Name())
					break
				}
			}
		}
		if strings.TrimSpace(intStr) == "" {
			intStr = "0"
		}

		intData, err = strconv.ParseInt(intStr, 10, 64)
		if err != nil {
			break
		}

		v.SetInt(intData)
	case reflect.Float32, reflect.Float64:
		if strings.TrimSpace(data) == "" {
			data = "0"
		}
		var floatData float64
		floatData, err = strconv.ParseFloat(data, 64)
		if err != nil {
			break
		}

		v.SetFloat(floatData)
	case reflect.Bool:
		var boolData bool
		boolData, err = strconv.ParseBool(data)
		if err != nil {
			break
		}

		v.SetBool(boolData)
	case reflect.String:
		v.SetString(data)
	default:
		value := "value " + data
		if data == "" {
			value = "empty value"
		}
		err = fmt.Errorf("can't set %s to type %s", value, v.Type().Name())
	}
	return
}
