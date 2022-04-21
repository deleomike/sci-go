package mean_shift

import (
	"testing"

	"github.com/deleomike/sci-go/internal"
)

func TestFit(t *testing.T) {
	var data = [][2]float64{
		{1, 2},
		{1.5, 1.8},
		{5, 8},
		{8, 8},
		{1, 0.6},
		{9, 11},
		{8, 2},
		{10, 2},
		{9, 3},
	}

	Fit(data, 4, 200000)

}

func TestPredict(t *testing.T) {
	var centroids = [][2]float64{
		{1.75, 2.1999999999999997},
		{8.5, 2.5},
	}

	var data = [][2]float64{
		{1, 2},
		{1.5, 1.8},
		{5, 8},
		{8, 8},
		{1, 0.6},
		{9, 11},
		{8, 2},
		{10, 2},
		{9, 3},
	}

	cluster_ids := Predict(centroids, data)

	real_cluster_ids := []int{0, 0, 1, 1, 0, 1, 1, 1, 1}

	if !internal.ArrayEqual[int](real_cluster_ids, cluster_ids) {
		t.Fail()
	}

}

func TestFitPredict(t *testing.T) {
	var data = [][2]float64{
		{1, 2},
		{1.5, 1.8},
		{5, 8},
		{8, 8},
		{1, 0.6},
		{9, 11},
		{8, 2},
		{10, 2},
		{9, 3},
	}

	real_cluster_ids := []int{0, 0, 1, 1, 0, 1, 1, 1, 1}

	cluster_ids := FitPredict(data, 4, 200000)

	if !internal.ArrayEqual[int](real_cluster_ids, cluster_ids) {
		t.Fail()
	}

}

func BenchmarkFit(b *testing.B) {

	// fmt.Println(fit(data, uint(2), uint(10)))

	var data = [][2]float64{
		{1, 2},
		{1.5, 1.8},
		{5, 8},
		{8, 8},
		{1, 0.6},
		{9, 11},
		{8, 2},
		{10, 2},
		{9, 3},
	}

	for i := 0; i < b.N; i++ {
		Fit(data, 3, 200)
	}

	// y = KMeans{0, 1}

	// fmt.Println(y)

}
