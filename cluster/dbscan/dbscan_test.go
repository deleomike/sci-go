package dbscan

import "testing"

func TestFitPredict(t *testing.T) {
	var data = [][2]float64{
		{1, 2},
		{1, 3},
		{2, 3.1},
		{2, 2.9},
		{1.9, 3.0},
		{2.1, 3.3},

		{1, 0},
		{10, 2},
		{10, 4},
		{10, 0},
	}

	FitPredict(data, 2, 1)
}
