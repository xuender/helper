package slice

import "github.com/xuender/helper/types"

// Reduce applies a reducer function to all elements of a slice and returns a single value.
// It iterates over the slice, applying the reducer to accumulate a result.
// The first element initializes the accumulator.
func Reduce[S ~[]T, T any](items S, reducer types.Reducer[T]) T {
	var ret T

	for idx, item := range items {
		if idx == 0 {
			ret = item
		} else {
			ret = reducer(ret, item)
		}
	}

	return ret
}
