package math

import "github.com/xuender/rg/types"

// Divide returns a function that divides the given dividend by the input divisor.
// The type parameter T must implement the Number interface.
func Divide[T types.Number](dividend T) types.Mapper[T, T] {
	return func(divisor T) T {
		return dividend / divisor
	}
}
