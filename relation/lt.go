package relation

import "cmp"

// Lt returns a function that checks if the input value is less than the given base value.
// The type parameter T must implement the cmp.Ordered interface.
func Lt[T cmp.Ordered](base T) func(T) bool {
	return func(val T) bool {
		return base < val
	}
}
