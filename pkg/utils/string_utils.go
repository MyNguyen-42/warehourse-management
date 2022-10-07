package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func FormatPhoneNumber(phoneNumber string) string {
	phoneNumber = strings.TrimSpace(strings.ToValidUTF8(phoneNumber, ""))
	if strings.HasPrefix(phoneNumber, "+84") {
		phoneNumber = "0" + strings.TrimLeft(phoneNumber, "+84")
	} else if strings.HasPrefix(phoneNumber, "84") {
		phoneNumber = "0" + strings.TrimLeft(phoneNumber, "84")
	} else if !strings.HasPrefix(phoneNumber, "0") {
		phoneNumber = "0" + strings.TrimLeft(phoneNumber, "+84")
	}

	return phoneNumber
}

// Compare compares two string pointers safety
func Compare(a, b *string) bool {
	if a == nil {
		return b == nil
	}

	return *a == *b
}

func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

var baseCode = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateCode(length int) string {
	return stringWithCharset(length, baseCode)
}

func NumberToCode(number int) string {
	value := ""
	base := len(baseCode)
	for number > 0 {
		index := number % base
		number = number / base
		value = string(baseCode[index]) + value
	}
	return value
}

func NewRandomCode(now time.Time, length int) string {
	dateCode := fmt.Sprintf("%s%s", NumberToCode(now.Year()-2000), NumberToCode(now.YearDay()))
	return dateCode + GenerateCode(length-len(dateCode))
}
