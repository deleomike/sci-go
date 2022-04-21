package main

import (
	"fmt"

	"github.com/deleomike/sci-go/cluster/kmeans"
)

func main() {

	var (
		data         [][2]float64
		centroids    [][2]float64
		cluster_ids  []int
		num_clusters uint
		max_iters    uint
	)

	// Dataset
	data = [][2]float64{
		{1, 2},
		{1, 4},
		{1, 0},
		{10, 2},
		{10, 4},
		{10, 0},
	}

	num_clusters = 3
	max_iters = 200

	// Fit a dataset

	centroids = kmeans.Fit(data, num_clusters, max_iters)
	fmt.Println(centroids)

	// Predict from Centroids

	cluster_ids = kmeans.Predict(centroids, data)
	fmt.Println(cluster_ids)

	// Fit Predict

	cluster_ids = kmeans.FitPredict(data, num_clusters, max_iters)
	fmt.Println(cluster_ids)

}
