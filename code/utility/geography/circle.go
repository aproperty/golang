package geography

// IsInCircleFence ...
func IsInCircleFence(lengh, circleLng, circleLat, lng, lat float64) bool {

	betweenMeter := EarthDistance(circleLat, circleLng, lat, lng)

	return betweenMeter < lengh
}
