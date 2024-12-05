package list

import "cmp"

// Min returns the smallest element from a slice of ordered elements.
//
// It iterates over the slice and compares each element to find the minimum value.
//
// If the slice is empty, it returns the zero value of the element type.
func Min[T cmp.Ordered](items []T) T {
	var ret T

	if len(items) == 0 {
		return ret
	}

	ret = items[0]

	for _, item := range items[1:] {
		if item < ret {
			ret = item
		}
	}

	return ret
}
