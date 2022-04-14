package scigo

import "math"

// EuclideanDistance is one of the common distance measurement
func EuclideanDistance(a, b []float64) float64 {
	var (
		s, t float64
	)

	for i, _ := range a {
		t = a[i] - b[i]
		s += t * t
	}

	return math.Sqrt(s)
}

// EuclideanDistanceSquared is one of the common distance measurement
func EuclideanDistanceSquared(a, b []float64) float64 {
	var (
		s, t float64
	)

	for i, _ := range a {
		t = a[i] - b[i]
		s += t * t
	}

	return s
}
