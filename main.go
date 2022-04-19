package main

import (
	"fmt"

	"github.com/deleomike/sci-go/cluster/kmeans"
)

func main() {

	fmt.Println(kmeans.Fit([][2]float64{}, 2, 2))

}
