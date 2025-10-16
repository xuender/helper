package fp

import (
	"github.com/xuender/helper/cont"
	"github.com/xuender/helper/gtype"
)

// All returns a function that checks if all elements in a cont satisfy all provided predicates.
func All[T any](chekders ...gtype.Checker[T]) func([]T) bool {
	return func(items []T) bool {
		return cont.All(items, chekders...)
	}
}
