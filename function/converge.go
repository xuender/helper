package function

import (
	"github.com/xuender/rg/list"
	"github.com/xuender/rg/types"
)

// Converge applies multiple functions to a slice and then reduces their results using a final function.
func Converge[T any](after types.Reducer[T], yields []types.ReducerSlice[T]) func(items []T) T {
	return func(items []T) T {
		data := make([]T, len(yields))
		for idx, yield := range yields {
			data[idx] = yield(items)
		}

		return list.Reduce(after)(data)
	}
}

// ConvergeCurry applies multiple functions to a slice and reduces their results using a curried final function.
func ConvergeCurry[T any](after types.ReducerCurry[T], yields []types.ReducerSlice[T]) func(items []T) T {
	return func(items []T) T {
		data := make([]T, len(yields))
		for idx, yield := range yields {
			data[idx] = yield(items)
		}

		return list.ReduceCurry(after)(data)
	}
}
