package slice

import "cmp"

func Min[T cmp.Ordered](items []T) T {
	var minVal T

	if len(items) == 0 {
		return minVal
	}

	minVal = items[0]

	for _, item := range items[1:] {
		if item < minVal {
			minVal = item
		}
	}

	return minVal
}
