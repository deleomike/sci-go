package scigo

import (
	"testing"
)

var data = [][2]float64{
	{1, 2},
	{1, 4},
	{1, 0},
	{10, 2},
	{10, 4},
	{10, 0},
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

	fit(data, 2, 100)
}

func BenchmarkXxx(b *testing.B) {

	// fmt.Println(fit(data, uint(2), uint(10)))

	for i := 0; i < b.N; i++ {
		fit_predict(data, uint(2), uint(1))
	}

	// y = KMeans{0, 1}

	// fmt.Println(y)

}
