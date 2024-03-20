package timeutil

import (
	"time"
)

func MustParseDuration(s string) time.Duration {
	value, err := time.ParseDuration(s)
	if err != nil {
		panic("util: Can't parse duration `" + s + "`: " + err.Error())
	}
	return value
}

func ParseInLocal(layout, value string) (time.Time, error) {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return time.Time{}, err
	}
	return time.ParseInLocation(layout, value, loc)
}

func GetTimestamp(value string) int64 {
	ti, _ := time.Parse(time.RFC3339, value)
	timeUint := ti.In(time.Local).Unix()
	return timeUint
}

func GetTimeFormat(str string) string {
	// ti, _ := time.Parse(time.RFC3339, str)
	ti, err := time.ParseInLocation("2006-01-02T15:04:05.000Z", str, time.Local)
	if err != nil {
		return ""
	}
	re := ti.Format("2006-01-02 15:04:05")
	return re
}
