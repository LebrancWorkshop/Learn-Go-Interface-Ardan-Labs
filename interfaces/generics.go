package interfaces

import (
	"fmt"
)

type number interface {
	~int | ~float64
}

func Max[T number](values []T) (T, error) {
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
