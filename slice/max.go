package slice

import "cmp"

func Max[T cmp.Ordered](items []T) T {
	var maxVal T

	if len(items) == 0 {
		return maxVal
	}

	maxVal = items[0]

	for _, item := range items[1:] {
		if item > maxVal {
			maxVal = item
		}
	}

	return maxVal
}
