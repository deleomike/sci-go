package dbscan

import (
	"fmt"

	"github.com/deleomike/sci-go/cluster"
)

func FitPredict(data [][2]float64, min_samples uint, eps float64) [][2]float64 {
	var (
		core_points     [][2]float64
		non_core_points [][2]float64
	)

	// Go through data points and identify core points

	//TODO theres an efficient way to go about this...right now it is brute force
	for i, point := range data {
		num_points := 0

		for j, pointB := range data {
			if i != j {
				distance := cluster.EuclidianDistance(point, pointB)

				if distance < eps {
					num_points++
				}
			}
		}

		if num_points >= int(min_samples) {
			core_points = append(core_points, point)
		} else {
			non_core_points = append(non_core_points, point)
		}
	}

	fmt.Println("Core Points", core_points)

	fmt.Println("Non Core Points", non_core_points)

	return core_points

}
