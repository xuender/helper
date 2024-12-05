package list

// Length returns the number of elements in a slice.
func Length[T any](items []T) int {
	return len(items)
}
