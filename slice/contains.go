package slice

import "slices"

// Contains returns a function that checks if a given slice contains the specified item.
func Contains[T comparable](item T) func([]T) bool {
	return func(items []T) bool {
		return slices.Contains(items, item)
	}
}
