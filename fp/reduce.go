package fp

import (
	"github.com/xuender/helper/cont"
	"github.com/xuender/helper/gtype"
)

// Reduce applies a function cumulatively to the elements of a cont.
func Reduce[T any](yield gtype.Reducer[T]) func(items []T) T {
	return func(items []T) T {
		return cont.Reduce(items, yield)
	}
}
