package cont

import "cmp"

// Min returns the minimum value from a slice of ordered elements.
// If the slice is empty, it returns the zero value of the element type.
func Min[S ~[]E, E cmp.Ordered](items S) E {
	var minVal E

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
