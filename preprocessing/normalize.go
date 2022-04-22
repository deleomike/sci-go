package preprocessing

import (
	"fmt"
	"github.com/deleomike/sci-go/internal"
)

// :param: X - 2D array
// :param: norm - normalization type {"l1", "l2", "max"}
func Normalize[T float32 | float64](X [][]T, norm string) [][]T{
	// TODO Error Handling

	var norm T

	switch norm {
	case "l1":
		norm = internal.Sum(internal.Abs[T](X))
		break
	case "l2":
		break
	case "max":
		break
	default:
		fmt.Println("Warning, incorrect normalization", norm, "given. Defaulting to Max. Options are 'l1', 'l2', 'max'")
		return Normalize[T](X, "max")
		break
	}

	return X \ norm


}
