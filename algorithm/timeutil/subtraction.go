package timeutil

import (
	"errors"
	"time"
)

// CalculateTimeDelay delay = two - one 单位是 s
func CalculateTimeDelay(one string, two string) (int64, error) {

	loc, ok := time.LoadLocation("Local")
	if ok != nil {
		return 0, ok
	}

	oneTime, ok1 := time.ParseInLocation(DATE_TIME, one, loc)
	if ok1 != nil {
		return 0, ok1
	}
	one1 := oneTime.Unix()

	twoTime, ok2 := time.ParseInLocation(DATE_TIME, two, loc)
	if ok2 != nil {
		return 0, ok2
	}
	two1 := twoTime.Unix()

	return two1 - one1, nil
}

// CalculateTimeDelayDay delay = two - one 单位是 day
func CalculateTimeDelayDay(one1 string, two1 string) (int64, error) {

	loc, ok := time.LoadLocation("Local")
	if ok != nil {
		return 0, ok
	}

	if len(one1) != 19 || len(two1) != 19 {
		return -1, errors.New("inputs error")
	}
	one := one1[:11] + "00:00:00"
	two := two1[:11] + "00:00:00"

	oneTime, ok1 := time.ParseInLocation(DATE_TIME, one, loc)
	if ok1 != nil {
		return 0, ok1
	}
	twoTime, ok2 := time.ParseInLocation(DATE_TIME, two, loc)
	if ok2 != nil {
		return 0, ok2
	}
	delays := int64(timeSubDays(twoTime, oneTime))

	if delays == -1 {
		return -1, errors.New("sub days error")
	}
	return delays, nil
}

func timeSubDays(t1, t2 time.Time) int {

	if t1.Location().String() != t2.Location().String() {
		return -1
	}

	hours := t1.Sub(t2).Hours()
	if hours <= 0 {
		return -1
	}

	if hours < 24 {
		t1y, t1m, t1d := t1.Date()
		t2y, t2m, t2d := t2.Date()
		isSameDay := (t1y == t2y && t1m == t2m && t1d == t2d)
		if isSameDay {
			return 0
		}
		return 1
	}

	if (hours/24)-float64(int(hours/24)) == 0 {
		return int(hours / 24)
	}

	return int(hours/24) + 1
}
