package mtsp

func pointIsOnTheLine(lineCoords [][]float64, pointCoords []float64) bool {
	//calculate slope of line
	slope := (lineCoords[1][1] - lineCoords[0][1]) / (lineCoords[1][0] - lineCoords[0][0])
	//cal the y Intercept of the line.
	yIntercept := lineCoords[0][1] - slope*lineCoords[0][0]
	//Calculate the y value of the point using the equation of the line.
	yValue := slope*pointCoords[0] + yIntercept
	return yValue == pointCoords[1]
}

// this Still need to optimize -- Algorithm can be better
// this function will make sure if the load2 route is already covered by load1 delivery without any extra cost
// i.e L1P----L2P-----L2D----------L1D///True else false
func DeliveryisSubset(load1, load2 Load) bool {
	lineCoords := [][]float64{{load1.Pickup.x, load1.Pickup.y}, {load1.Dropoff.x, load1.Dropoff.y}} //this will be delivery i
	point1Coords := []float64{load2.Pickup.x, load2.Pickup.y}
	point2Coords := []float64{load2.Dropoff.x, load2.Dropoff.y}
	if pointIsOnTheLine(lineCoords, point1Coords) && pointIsOnTheLine(lineCoords, point2Coords) {
		//pickup2
		loadD1 := Load{LoadNumber: "d1", Pickup: coords{x: load1.Pickup.x, y: load1.Pickup.y}, Dropoff: coords{x: load1.Dropoff.x, y: load1.Dropoff.y}}
		d1 := calculateLoadDistance(loadD1) //pickup1->dropoff1

		loadD2 := Load{LoadNumber: "d2", Pickup: coords{x: load1.Pickup.x, y: load1.Pickup.y}, Dropoff: coords{x: load2.Pickup.x, y: load2.Pickup.y}}
		d2 := calculateLoadDistance(loadD2) //pickup1->pickup2

		loadD3 := Load{LoadNumber: "d3", Pickup: coords{x: load1.Dropoff.x, y: load1.Dropoff.y}, Dropoff: coords{x: load2.Pickup.x, y: load2.Pickup.y}}
		d3 := calculateLoadDistance(loadD3) //dropoff1->pickup2
		if d2 < d1 && d3 < d1 {
			//dropoff2
			loadD2 = Load{LoadNumber: "d2", Pickup: coords{x: load1.Pickup.x, y: load1.Pickup.y}, Dropoff: coords{x: load2.Dropoff.x, y: load2.Dropoff.y}}
			d2 = calculateLoadDistance(loadD2) //pickup1->dropof2
			loadD3 = Load{LoadNumber: "d3", Pickup: coords{x: load1.Dropoff.x, y: load1.Dropoff.y}, Dropoff: coords{x: load2.Dropoff.x, y: load2.Dropoff.y}}
			d3 = calculateLoadDistance(loadD3) //dropoff1->dropoff2
			if d2 < d1 && d3 < d1 {
				return true
			}
		}
	}
	return false
}
