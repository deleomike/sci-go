package internal

import (
	"math"
)

// NOTE: this file is really just a placeholder for certain functions. Most of this
// will be replaced with external functionality

func Abs[T int | float32 | float64](x [][]T) [][]T {

	//TODO error checking
	var res [][]T

	res = x

	for i := 0; i < len(res); i++{
		for j := 0; j < len(res[0]); j++ {
			res[i][j] = math.Abs(res[i][j])
		}
	}

	return res
}

func Sum[T int | float32 | float64](x [][]T) T {
	var sum T

	for i := 0; i < len(x); i++{
		for j := 0; j < len(x[0]); j++ {
			sum += x[i][j]
		}
	}

	return sum
}