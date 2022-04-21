package internal

import "fmt"

func ArrayEqual[T int, float32, float64, string] (a, b []T) bool{
	fmt.Println(a)
	fmt.Println(b)
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}