package kmeans

import (
	"math"
	"testing"
)

func intArrayEquals(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestFit(t *testing.T) {
	//todo

	var data = [][2]float64{
		{1, 2},
		{1, 4},
		{1, 0},
		{10, 2},
		{10, 4},
		{10, 0},
	}

	var centroids = Fit(data, 2, 100)

	for _, c := range centroids {
		if math.IsNaN(c[0]) || math.IsNaN(c[1]) {
			t.FailNow()
		}
	}
}

func TestPredict(t *testing.T) {
	var data = [][2]float64{
		{1, 2},
		{1, 4},
		{1, 0},
		{10, 2},
		{10, 4},
		{10, 0},
	}

	var centroids = [][2]float64{
		{1.5, 3},
		{15, 3},
	}

	var cluster_ids = Predict(centroids, data)

	if !(intArrayEquals(cluster_ids, []int{0, 0, 0, 1, 1, 1}) ||
		intArrayEquals(cluster_ids, []int{1, 1, 1, 0, 0, 0})) {

		t.FailNow()

	}

}

func TestFitPredict(t *testing.T) {
	var data = [][2]float64{
		{1, 2},
		{1, 4},
		{1, 0},
		{10, 2},
		{10, 4},
		{10, 0},
	}

	var cluster_ids = FitPredict(data, 2, 100)

	if !(intArrayEquals(cluster_ids, []int{0, 0, 0, 1, 1, 1}) ||
		intArrayEquals(cluster_ids, []int{1, 1, 1, 0, 0, 0})) {

		t.FailNow()

	}
}

func BenchmarkFitPredict(b *testing.B) {

	// fmt.Println(fit(data, uint(2), uint(10)))

	var data = [][2]float64{
		{1, 2},
		{1, 4},
		{1, 0},
		{10, 2},
		{10, 4},
		{10, 0},
	}

	for i := 0; i < b.N; i++ {
		FitPredict(data, uint(2), uint(1))
	}

	// y = KMeans{0, 1}

	// fmt.Println(y)

}
