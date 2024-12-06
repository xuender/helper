package slice

import "github.com/xuender/helper/types"

func Reduce[T any](items []T, reducer types.Reducer[T]) T {
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
