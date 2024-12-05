package math

import "github.com/xuender/rg/types"

// Multiply returns a function that multiplies a given value by a base.
func Multiply[T types.Number](base T) func(T) T {
	return func(val T) T {
		return base * val
	}
}
