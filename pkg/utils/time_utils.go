package utils

import (
	"fmt"
	"strconv"
	"time"
)

func DateIndex(t time.Time) int {
	value64, _ := strconv.ParseInt(t.Format("20060102"), 10, 64)
	return int(value64)
}

//func TwoThousandYearsLater() time.Time {
//	return time.Date(4022, 1, 1, 0, 0, 0, 0, constants.CurrentLoc())
//}

func TimeStartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func TimeIndex(t time.Time) int {
	return t.Hour()*10000 + t.Minute()*100 + t.Second()
}

func SerializeTime(t time.Time) string {
	return fmt.Sprintf("%04d-%02d-%02dT%02d:%02d:%02d.%03dZ", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(),
		t.Second(), t.Nanosecond()/1000000)
}

func DeserializeTime(t string) (*time.Time, error) {
	val, err := time.Parse(time.RFC3339, t)
	if err != nil {
		return nil, err
	}
	return &val, err
}

func UnixMilli(msec int64) time.Time {
	return time.Unix(msec/1e3, (msec%1e3)*1e6)
}

func ParseTimeWithFormat(s string, format string) (time.Time, error) {
	return time.Parse(format, s)
}

func Weekdays(weekdays []string) []time.Weekday {
	var rs []time.Weekday
	for _, weekday := range weekdays {
		rs = append(rs, Weekday(weekday))
	}
	return rs
}

func Weekday(weekday string) time.Weekday {
	switch weekday {
	case time.Monday.String():
		return time.Monday
	case time.Tuesday.String():
		return time.Tuesday
	case time.Wednesday.String():
		return time.Wednesday
	case time.Thursday.String():
		return time.Thursday
	case time.Friday.String():
		return time.Friday
	case time.Saturday.String():
		return time.Saturday
	default:
		return time.Sunday
	}
}

func WeekDayVietnamese(data string) time.Weekday {
	switch data {
	case "THỨ 2":
		return time.Weekday(1)
	case "THỨ 3":
		return time.Weekday(2)
	case "THỨ 4":
		return time.Weekday(3)
	case "THỨ 5":
		return time.Weekday(4)
	case "THỨ 6":
		return time.Weekday(5)
	case "THỨ 7":
		return time.Weekday(6)
	case "CHỦ NHẬT":
		return time.Weekday(0)
	default:
		return time.Now().Weekday()
	}
}

//func HourToTime(now time.Time, hour string) (time.Time, error) {
//	timeHour, err := time.Parse("15:04", hour)
//	if err != nil {
//		return now, err
//	}
//	return time.Date(now.Year(), now.Month(), now.Day(), timeHour.Hour(), timeHour.Minute(), 0, 0, constants.CurrentLoc()), nil
//}
//
//func StartDayOf(which time.Time) time.Time {
//	return time.Date(which.Year(), which.Month(), which.Day(), 0, 0, 0, 0, constants.CurrentLoc())
//}
//
//func EndDayOf(which time.Time) time.Time {
//	return time.Date(which.Year(), which.Month(), which.Day(), 23, 59, 59, 999, constants.CurrentLoc())
//}

func ConvertWeekDayToVietnamese(data time.Weekday) string {
	switch data {
	case time.Weekday(1):
		return "THỨ 2"
	case time.Weekday(2):
		return "THỨ 3"
	case time.Weekday(3):
		return "THỨ 4"
	case time.Weekday(4):
		return "THỨ 5"
	case time.Weekday(5):
		return "THỨ 6"
	case time.Weekday(6):
		return "THỨ 7"
	case time.Weekday(0):
		return "CHỦ NHẬT"
	default:
		return ""
	}
}
