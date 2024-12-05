package function

import "github.com/xuender/rg/types"

// CurryReducer converts a binary reducer into a curried mapper.
func CurryReducer[T any](reducer types.Reducer[T]) types.ReducerCurry[T] {
	return func(val1 T) types.Mapper[T, T] {
		return func(val2 T) T {
			return reducer(val1, val2)
		}
	}
}
