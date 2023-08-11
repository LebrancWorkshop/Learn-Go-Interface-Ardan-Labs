package interfaces

import (
	"fmt"
)

func MaxInts(values []int) (int, error) {
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
