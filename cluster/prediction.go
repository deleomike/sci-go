package cluster

import "math"

func EuclidianDistance(pointA, pointB [2]float64) float64 {
	var dx = pointB[0] - pointA[0]
	var dy = pointB[1] - pointA[1]

	return math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))
}

func Predict(centroids [][2]float64, data [][2]float64) []int {

	clusters := make([]int, len(data))

	for i, x := range data {
		// Find each nearest centroid cj
		var min_distance float64 = math.Inf(1)
		var min_index int
		for j, c := range centroids {
			var _distance = EuclidianDistance(x, c)

			if math.Min(min_distance, _distance) == _distance {
				min_index = j
				min_distance = _distance
			}
		}

		// Assign the point xi to cluster j
		clusters[i] = min_index
	}

	return clusters

}
