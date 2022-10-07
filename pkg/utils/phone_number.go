package utils

import (
	"encoding/binary"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"regexp"
	"strings"
	"time"
)

const weirdChar = "\uFFFD"

func ToValidUTF8(s string) string {
	return strings.TrimSpace(strings.ToValidUTF8(s, weirdChar))
}

func ToValidUTF8s(ss []string) (rs []string) {
	for _, s := range ss {
		rs = append(rs, strings.TrimSpace(strings.ToValidUTF8(s, weirdChar)))
	}
	return
}

func IsValidOrderCode(orderCode string) bool {
	if len(orderCode) == 0 {
		return false
	}
	var validId = regexp.MustCompile("^[A-Za-z0-9_.-]*$")
	code := orderCode
	//if strings.HasSuffix(code, "_PR") {
	//	code = code[:len(code) - 3]
	//}
	result := validId.MatchString(code)
	return result
}

func InStringSlice(check string, ss []string) bool {
	for _, s := range ss {
		if s == check {
			return true
		}
	}
	return false
}

func GenerateId(id primitive.ObjectID) string {
	keyString := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(id.Hex())
	var number []byte
	number = append(number, 0)
	for i := 0; i < 6; i += 2 {
		b := bytes[len(bytes)-6+i]*16 + bytes[len(bytes)-6+i+1]
		number = append(number, b)
	}
	counter := binary.BigEndian.Uint32(number)

	now := time.Now()
	calTime := now.Sub(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))
	ms := calTime.Microseconds()
	ms = ms << 2
	ms = ms | int64(counter)
	if ms < 0 {
		ms = -ms
	}

	value := ""
	for ms > 0 {
		mod := ms % int64(len(keyString))
		ms = ms / int64(len(keyString))
		value = string(keyString[int(mod)]) + value
	}
	return value
}

func ConvertPhoneNumber(phone string) string {
	if len(phone) < 8 || len(phone) > 20 {
		return phone
	}

	phone = strings.Trim(phone, "\t \n")

	var prefix = phone[0:1]
	if prefix == "0" {
		phone = phone[1:]
		phone = "84" + phone
	}

	prefix = phone[0:3]
	if prefix == "+84" {
		phone = phone[1:]
	}

	return phone
}
