package fp

import (
	"github.com/xuender/helper/cont"
	"github.com/xuender/helper/gtype"
)

// Any returns a function that checks if any element in a cont satisfies any provided predicate.
func Any[T any](checkers ...gtype.Checker[T]) func([]T) bool {
	return func(items []T) bool {
		return cont.Any(items, checkers...)
	}
}
