package math

import "github.com/xuender/rg/types"

// Subtract returns a function that subtracts the input value from the given base value.
// The type parameter T must implement the Number interface.
func Subtract[T types.Number](base T) func(T) T {
	return func(val T) T {
		return base - val
	}
}
