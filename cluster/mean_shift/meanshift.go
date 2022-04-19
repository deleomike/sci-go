package mean_shift

import (
	"math"
)

func EstimateBandwidth() {

}

func sum(array [][2]float64) [2]float64 {
	var sum = [2]float64{0.0, 0.0}

	for _, val := range array {
		sum[0] += val[0]
		sum[1] += val[1]
	}

	return sum
}

func average(array [][2]float64) [2]float64 {

	var sum = sum(array)

	return [2]float64{
		sum[0] / float64(len(sum)),
		sum[1] / float64(len(sum)),
	}
}

func norm(a, b [2]float64) float64 {
	return math.Sqrt(math.Pow(a[0]-b[0], 2) + math.Pow(a[1]-b[1], 2))
}

func arr_equal(a, b [2]float64) bool {
	return int(a[0]*1000.0) == int(b[0]*1000.0) && int(a[1]*1000.0) == int(b[1]*1000.0)
}

// func bubble_sort(array [][2]float64) [][2]float64 {
// 	var res := [][2]float64 {}

// 	for i := 0; i < arr
// }

func Fit(data [][2]float64, bandwidth float64, max_iters uint) [][2]float64 {

	var (
		centroids      [][2]float64
		new_centroids  [][2]float64
		prev_centroids [][2]float64
		in_bandwidth   [][2]float64
		unique         [][2]float64
		done           bool
	)

	centroids = data

	for iter := 0; iter < int(max_iters); iter++ {
		new_centroids = [][2]float64{}

		for _, centroid := range centroids {
			in_bandwidth = [][2]float64{}

			for _, feature_set := range data {
				//TODO Norm
				if norm(feature_set, centroid) < bandwidth {
					in_bandwidth = append(in_bandwidth, feature_set)
				}
			}

			new_centroids = append(new_centroids, average(in_bandwidth))
		}

		// fmt.Println("new centroids", new_centroids)

		// TODO Do Unique Set...im on a plane, I dont have internet to download golang set
		// TODO sort the set
		unique = [][2]float64{}
		for i, _ := range new_centroids {
			var exists bool
			exists = false

			for j, _ := range unique {
				if arr_equal(unique[j], new_centroids[i]) {
					exists = true
					break
				}
			}

			if !exists {
				unique = append(unique, new_centroids[i])
			}
		}

		// unique.so

		prev_centroids = centroids

		// fmt.Println("unique", unique)
		// fmt.Println("prev centroids", prev_centroids)

		// fmt.Println(new_centroids)
		// fmt.Println(prev_centroids)

		done = true

		for i, _ := range unique {
			if !arr_equal(unique[i], prev_centroids[i]) {
				// fmt.Println("not equal", i)
				done = false
				// TODO. Maybe a candidate for GOTO
				break
			}
			// fmt.Println("these are equal?", unique[i], prev_centroids[i])
		}

		centroids = unique

		if done {
			break
		}

		// fmt.Println("")

	}

	return centroids

}
