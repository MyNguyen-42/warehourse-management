package zapfield

import (
	"encoding/json"
	"github.com/segmentio/ksuid"
	"go.uber.org/zap"
	"reflect"
	"time"
)

// jsonString returns json string marshal from objects
func jsonString(key string, value interface{}) zap.Field {
	marshal, err := json.Marshal(value)
	if err != nil {
		panic("jsonString got value not a json")
	}
	return zap.String(key, string(marshal))
}

func RequestId() zap.Field {
	return zap.String("request_id", ksuid.New().String())
}

func Method(reqData string) zap.Field {
	return zap.String("method", reqData)
}

func RequestData(reqData interface{}) zap.Field {
	return jsonString("request_data", reqData)
}

func RequestDataBin(data []byte) zap.Field {
	return zap.Binary("request_data", data)
}

func RequestIdentify(id string) zap.Field {
	return zap.String("request_id", id)
}

func RequestHub(hubId string) zap.Field {
	return zap.String("request_hub", hubId)
}

func ResponseData(resData interface{}) zap.Field {
	return jsonString("response_data", resData)
}

func GrpcAppCode(appCode string) zap.Field {
	return zap.String("grpc_app_code", appCode)
}

func HttpPath(path string) zap.Field {
	return zap.String("http_path", path)
}

func ExecutionTime(duration time.Duration) zap.Field {
	return zap.Int64("execution_time", duration.Milliseconds())
}

func OpsRequest(path string, request, response interface{}, duration time.Duration) []zap.Field {
	if response != nil && reflect.Indirect(reflect.ValueOf(response)).IsZero() {
		response = nil
	}
	return []zap.Field{
		zap.Int64("execution_time", duration.Milliseconds()),
		zap.String("http_path", path),
		RequestData(request),
		ResponseData(response),
	}
}

func ConsumerId(id string) zap.Field {
	return zap.String("consumer_id", id)
}

func CampaignCode(code string) zap.Field {
	return zap.String("campaign_code", code)
}

func QueueName(name string) zap.Field {
	return zap.String("queue_name", name)
}

func RemoteAddr(data string) zap.Field {
	return zap.String("remote_addr", data)
}

func MessageKey(key string) zap.Field {
	return zap.String("message_key", key)
}
