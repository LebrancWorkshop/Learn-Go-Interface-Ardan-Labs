package interfaces

import (
	"fmt"
)

func Max[T int | float64](values []T) (T, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("empty slice")
	}

	max := values[0]
	for _, v := range values[1:] {
		if v > max {
			max = v
		}
	}

	return max, nil
}
