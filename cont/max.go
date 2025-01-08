package cont

import "cmp"

// Max returns the maximum value from a slice of ordered elements.
// If the slice is empty, it returns the zero value of the element type.
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
