package utils

import (
	"encoding/json"
	"strconv"
)

// ParseInt convert string to int
func ParseInt(text string, defaultValue int) int {
	if text == "" {
		return defaultValue
	}

	num, err := strconv.Atoi(text)
	if err != nil {
		return defaultValue
	}
	return num
}

func ParseInt64FromInterface(input interface{}, defaultValue int64) int64 {
	bytes, err := json.Marshal(input)
	if err != nil {
		return defaultValue
	}

	var val int64
	err = json.Unmarshal(bytes, &val)
	if err != nil {
		return defaultValue
	}

	return val
}

func ToInt64(i interface{}) int64 {
	switch i.(type) {
	case int:
		return int64(i.(int))
	case int8:
		return int64(i.(int8))
	case int16:
		return int64(i.(int16))
	case int32:
		return int64(i.(int32))
	case int64:
		return i.(int64)
	case uint:
		return int64(i.(uint))
	case uint8:
		return int64(i.(uint8))
	case uint16:
		return int64(i.(uint16))
	case uint32:
		return int64(i.(uint32))
	case uint64:
		return int64(i.(uint64))
	case float32:
		return int64(i.(float32))
	case float64:
		return int64(i.(float64))
	default:
		return 0
	}
}

func ToInt32(i interface{}) int32 {
	switch i.(type) {
	case int:
		return int32(i.(int))
	case int8:
		return int32(i.(int8))
	case int16:
		return int32(i.(int16))
	case int32:
		return i.(int32)
	case int64:
		return int32(i.(int64))
	case uint:
		return int32(i.(uint))
	case uint8:
		return int32(i.(uint8))
	case uint16:
		return int32(i.(uint16))
	case uint32:
		return int32(i.(uint32))
	case uint64:
		return int32(i.(uint64))
	case float32:
		return int32(i.(float32))
	case float64:
		return int32(i.(float64))
	default:
		return 0
	}
}

// ParseInt64 convert string to int64
func ParseInt64(text string, defaultValue int64) int64 {
	if text == "" {
		return defaultValue
	}

	num, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		return defaultValue
	}
	return num
}
