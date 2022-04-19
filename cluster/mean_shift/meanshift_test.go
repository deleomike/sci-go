package mean_shift

import (
	"fmt"
	"testing"
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

	fmt.Println(Fit(data, 3, 200))

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
