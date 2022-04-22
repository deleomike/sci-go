package preprocessing

import (
	"fmt"
	"testing"
)

func SimpleNormalizeL1(t *testing.T) {
	x := [][]float32{
		{1, -1, 2},
		{2, 0, 0},
		{0, 1, -1},
	}

	fmt.Println(Normalize[float32](x, "l1"))
}
