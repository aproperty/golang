package fence

func IsInCircleFence(circleRadius, centerLng, centerLat,
	lng, lat float64) bool {
	betweenMeter := EarthDistance(centerLat, centerLng, lat, lng)
	return betweenMeter < circleRadius
}

// http://alienryderflex.com/polygon/
func IsInPolygon(polyX []float64, polyY []float64,
	x float64, y float64) bool {

	var i, j int
	i = 0
	j = len(polyX) - 1

	var oddNodes = false

	for i = 0; i < len(polyX); i++ {

		if ((polyY[i] < y && polyY[j] >= y) ||
			(polyY[i] >= y && polyY[j] < y)) &&
			(polyX[i] <= x || polyX[j] <= x) {

			if polyX[i]+
				(y-polyY[i])/(polyY[j]-polyY[i])*
					(polyX[j]-polyX[i]) < x {
				oddNodes = !oddNodes
			}
		}

		j = i
	}

	return oddNodes
}
