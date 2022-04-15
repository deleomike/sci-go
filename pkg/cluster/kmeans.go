package scigo

import (
	"math"
)

// Importing fmt package for the sake of printing

// Animal is the name we want but since we are

type KMeaner interface {
	fit()
	predict()
	fit_predict()
}

// func euclidian_distance(setA, setB []float32) []int {

// 	val := []int{}

// 	return val

// }

func distance(pointA, pointB [2]float64) float64 {
	var dx = pointB[0] - pointA[0]
	var dy = pointB[1] - pointA[1]

	return math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))
}

func Fit(data [][2]float64, num_clusters uint, max_iters uint) [][2]float64 {
	// if km.num_clusters > len(dataset) {
	// 	// TODO, some 'error' checking
	// }

	// INPUT x 1d array

	// Place centroids c1...ck at random locations
	centroids := make([][2]float64, num_clusters)

	// fmt.Println(centroids)

	// Init the clusters for the points
	clusters_new := make([]int, len(data))

	// fmt.Println(clusters_new)

	// REPEAT BELOW TILL CONVERGENCE

	for iterations := 0; iterations < int(max_iters); iterations++ {
		// clusters_old := clusters_new
		cluster_counts := make([]float64, num_clusters)

		// For each point xi
		for i, x := range data {
			// Find each nearest centroid cj
			var min_distance float64 = math.Inf(1)
			var min_index int
			for j, c := range centroids {
				var _distance = distance(x, c)

				if math.Min(min_distance, _distance) == _distance {
					min_index = j
					min_distance = _distance
				}
			}

			// Assign the point xi to cluster j
			clusters_new[i] = min_index
			cluster_counts[min_index] += 1

		}

		// fmt.Println(clusters_new)
		// fmt.Println(centroids)

		// For each cluster j = 1..k
		for i, x := range data {

			j := clusters_new[i]

			// New centroid cj = mean of all points assigned to cluster j
			centroids[j][0] += x[0]
			centroids[j][1] += x[1]
		}
		for j, count := range cluster_counts {
			if count == 0 {
				continue
			}
			centroids[j][0] /= count
			centroids[j][1] /= count
		}

		// fmt.Println(centroids)

		// STOP when no cluster assignments dont change
	}

	// fmt.Println(clusters_new)

	return centroids

}

func Predict(centroids [][2]float64, data [][2]float64) []int {

	clusters := make([]int, len(data))

	for i, x := range data {
		// Find each nearest centroid cj
		var min_distance float64 = math.Inf(1)
		var min_index int
		for j, c := range centroids {
			var _distance = distance(x, c)

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

func FitPredict(data [][2]float64, num_clusters uint, max_iters uint) []int {
	km := Fit(data, num_clusters, max_iters)

	predictions := Predict(km, data)

	return predictions
}
