package math

import (
	"cmp"

	"github.com/xuender/rg/types"
)

// Add returns a function that adds the input value to the given base value.
// The type parameter T must implement the cmp.Ordered interface.
func Add[T cmp.Ordered](base T) types.Mapper[T, T] {
	return func(val T) T {
		return base + val
	}
}
