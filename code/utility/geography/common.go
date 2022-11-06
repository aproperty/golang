package geography

import (
	"strings"
)

// LngRemovePrefix ...
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

// LatRemovePrefix ...
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
