package fp

import (
	"github.com/xuender/helper/slice"
	"github.com/xuender/helper/types"
)

// None returns a function that checks if no elements in a slice satisfy any provided predicate.
func None[T any](checkers ...types.Checker[T]) func([]T) bool {
	return func(items []T) bool {
		return slice.None(items, checkers...)
	}
}
