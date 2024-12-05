package relation

import "github.com/xuender/rg/types"

// Equals returns a function that checks if the input value is equal to the given base value.
// The type parameter T must implement the comparable interface.
func Equals[T comparable](base T) types.Checker[T] {
	return func(val T) bool {
		return val == base
	}
}
