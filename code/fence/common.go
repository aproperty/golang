package fence

import (
	"fmt"
	"strconv"
	"strings"
)

// func LngRemovePrefix(lng1 string) string {
// 	_lng1 := ""
// 	if lng1[0] == 'W' {
// 		_lng1 = "-" + lng1[1:]

// 	} else if lng1[0] == 'E' {
// 		_lng1 = lng1[1:]
// 	} else {
// 		_lng1 = lng1[:]
// 	}
// 	return _lng1
// }

// func LatRemovePrefix(lat1 string) string {
// 	_lat1 := ""
// 	if lat1[0] == 'S' {
// 		_lat1 = "-" + lat1[1:]

// 	} else if lat1[0] == 'N' {
// 		_lat1 = lat1[1:]
// 	} else {
// 		_lat1 = lat1[:]
// 	}
// 	return _lat1
// }

func stringNotNull(str string) bool {
	if str == "" {
		return false
	} else if strings.EqualFold(str, "-") {
		return false
	} else if strings.EqualFold(str[1:], "0.000000") {
		return false
	} else {
		return true
	}
}

// exchangeXY 经纬度转换成 float64 类型
func exchangeXY(lat1 string, lng1 string) (la float64, ln float64, err bool) {

	_lat1 := ""
	if lat1[0] == 'S' {
		_lat1 = "-" + lat1[1:]
	} else if lat1[0] == 'N' {
		_lat1 = lat1[1:]
	} else {
		_lat1 = lat1[:]
	}
	var ok error
	la, ok = strconv.ParseFloat(_lat1, 64)
	if ok != nil {
		return 0, 0, false
	}

	_lng1 := ""
	if lng1[0] == 'W' {
		_lng1 = "-" + lng1[1:]
	} else if lng1[0] == 'E' {
		_lng1 = lng1[1:]
	} else {
		_lng1 = lng1[:]
	}
	ln, ok = strconv.ParseFloat(_lng1, 64)
	if ok != nil {
		return 0, 0, false
	}

	fmt.Println("Latitude=" + strconv.FormatFloat(la, 'f', -1, 32) + ";longitude=" + strconv.FormatFloat(ln, 'f', -1, 32))
	return la, ln, true
}
