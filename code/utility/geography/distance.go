package geography

import (
	"fmt"
	"math"
	"strconv"
)

// EarthDistanceString ...
func EarthDistanceString(lat1, lng1, lat2, lng2 string) float64 {

	if stringNotNull(lat1) && stringNotNull(lng1) && stringNotNull(lat2) && stringNotNull(lng2) {

		lat11, lng11, _ := exchangeXY(lat1, lng1)
		lat22, lng22, _ := exchangeXY(lat2, lng2)

		dis := EarthDistance(lat11, lng11, lat22, lng22)

		if math.IsNaN(dis) {
			return 0.0
		}
		return dis
	}

	return 0.0
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

// EarthDistance
// unit of return value is meter
func EarthDistance(lat1, lng1, lat2, lng2 float64) float64 {

	radius := 6378137 // W84 坐标基数 // 6371000

	rad := math.Pi / 180.0

	lat1 = lat1 * rad
	lng1 = lng1 * rad
	lat2 = lat2 * rad
	lng2 = lng2 * rad

	theta := lng2 - lng1
	dist := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))

	return dist * float64(radius)
}
