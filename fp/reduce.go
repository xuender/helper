package fp

import (
	"github.com/xuender/helper/slice"
	"github.com/xuender/helper/types"
)

// Reduce applies a function cumulatively to the elements of a slice.
func Reduce[T any](yield types.Reducer[T]) func(items []T) T {
	return func(items []T) T {
		return slice.Reduce(items, yield)
	}
}
