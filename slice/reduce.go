package slice

import "github.com/xuender/helper/types"

// Reduce applies a function cumulatively to the elements of a slice.
func Reduce[T any](yield types.Reducer[T]) func(items []T) T {
	return func(items []T) T {
		var ret T

		for idx, item := range items {
			if idx == 0 {
				ret = item
			} else {
				ret = yield(ret, item)
			}
		}

		return ret
	}
}
