package utils

func ConvertModels[Model interface{}, DTO interface{}](models []Model, fun func(data Model) DTO) []DTO {
	var rs []DTO
	for _, data := range models {
		rs = append(rs, fun(data))
	}
	return rs
}

func DataToGroupMap[Model interface{}](models []Model, fun func(data Model) string) map[string][]Model {
	var dataMap = make(map[string][]Model)
	for _, data := range models {
		key := fun(data)
		dataMap[key] = append(dataMap[key], data)
	}
	return dataMap
}

func GetKeys[M any, T any](models []M, fun func(data M) T) []T {
	var keys []T
	for _, data := range models {
		keys = append(keys, fun(data))
	}
	return keys
}
