package relation

import "cmp"

// Gt returns a function that checks if the input value is greater than the given base value.
// The type parameter T must implement the cmp.Ordered interface.
func Gt[T cmp.Ordered](base T) func(T) bool {
	return func(val T) bool {
		return base > val
	}
}
