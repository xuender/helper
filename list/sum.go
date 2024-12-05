package list

import "cmp"

// Sum calculates the sum of elements in a slice.
func Sum[T cmp.Ordered](items []T) T {
	var ret T

	for _, item := range items {
		ret += item
	}

	return ret
}
