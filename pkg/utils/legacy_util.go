package utils

import (
	"encoding/json"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"hash/fnv"
	"strings"
	"time"
)

func Hash(s string) (uint64, error) {
	h := fnv.New64a()
	_, err := h.Write([]byte(s))
	if err != nil {
		return 0, err
	}

	return h.Sum64(), nil
}

// to struct
func FromJson(data []byte, v interface{}) error {
	// call UnMarshal
	err := json.Unmarshal(data, v)
	return err
}

func FromBson(data []byte, v interface{}) error {
	// call UnMarshal
	err := bson.UnmarshalExtJSON(data, true, v)
	return err
}

// to json
func ToJson(object interface{}) ([]byte, error) {
	// call Marshal
	b, err := json.Marshal(&object)
	return b, err
}

func Contains(arr []string, key string) bool {

	if arr == nil || len(arr) == 0 {
		return false
	}

	for i := range arr {
		if key == arr[i] {
			return true
		}
	}

	return false
}

func ParseTime(str string) (time.Time, error) {

	if len(str) == 0 {
		return time.Now(), errors.New("time input nil")
	}

	t, e := parseTime(str)
	if e == nil {
		return t, nil
	}

	return time.Now(), e
}

func parseTime(str string) (time.Time, error) {
	layout := "2006-01-02T15:04:05Z07:00"
	t, err := time.Parse(layout, str)
	return t, err
}

func Remove(arr []string, key string) []string {

	index := -1

	for i, v := range arr {
		if v == key {
			index = i
			break
		}
	}

	if index != -1 {
		arr = append(arr[:index], arr[index+1:]...)
	}

	return arr
}

func FormatYYYYMMDD_HHmmss(t *time.Time) string {
	if t == nil || t.IsZero() {
		return ""
	}
	value := t.Format("2006-01-02 15:04:05")
	value = strings.ReplaceAll(value, ":", "")
	value = strings.ReplaceAll(value, "-", "")
	arr := strings.Split(value, " ")
	return arr[0] +"_"+arr[1]
}

func FormatYYYYMMDD(t *time.Time) string {
	if t == nil || t.IsZero() {
		return ""
	}
	value := t.Format("2006-01-02 15:04:05")
	value = strings.ReplaceAll(value, "-", "")
	arr := strings.Split(value, " ")
	return arr[0]

}
