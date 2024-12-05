package list

import "github.com/xuender/rg/types"

func ZipWith[T any](reducer types.Reducer[T], items1 []T) func([]T) []T {
	return func(items2 []T) []T {
		length := Min([]int{len(items1), len(items2)})
		ret := make([]T, length)

		for idx := range length {
			ret[idx] = reducer(items1[idx], items2[idx])
		}

		return ret
	}
}
