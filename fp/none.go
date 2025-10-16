package fp

import (
	"github.com/xuender/helper/cont"
	"github.com/xuender/helper/gtype"
)

// None returns a function that checks if no elements in a cont satisfy any provided predicate.
func None[T any](checkers ...gtype.Checker[T]) func([]T) bool {
	return func(items []T) bool {
		return cont.None(items, checkers...)
	}
}
