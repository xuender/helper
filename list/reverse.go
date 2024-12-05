package list

import "slices"

// Reverse reverses the order of elements in a slice.
func Reverse[T any](items []T) []T {
	slices.Reverse(items)

	return items
}
