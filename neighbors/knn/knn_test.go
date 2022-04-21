package knn

import "testing"

func TestFit(t *testing.T) {
	var data = [][2]float64{
		{1, 2},
		{1, 4},
		{1, 0},
		{10, 2},
		{10, 4},
		{10, 0},
	}

	Fit(data, 2)
}
