package knn

import (
	"fmt"
	"math"
)

func distance(pointA, pointB [2]float64) float64 {
	var dx = pointB[0] - pointA[0]
	var dy = pointB[1] - pointA[1]

	return math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))
}

func Fit(data [][2]float64, k uint) {

	for i, p1 := range data {
		distances := []float64{}
		for j, p2 := range data {
			if i != j {
				distances = append(distances, distance(p1, p2))
			}
		}
		// TODO Sort

		fmt.Println(distances)
	}

}
