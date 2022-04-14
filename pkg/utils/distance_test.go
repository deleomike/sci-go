package scigo

import (
	"testing"
)

func BenchmarkEuclidianDistance(b *testing.B) {

	var a = []float64{0, 1, 2, 100, 200}

	for i := 0; i < b.N; i++ {
		EuclideanDistance(a, a)
	}
}
