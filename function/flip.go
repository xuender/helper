package function

import "github.com/xuender/rg/types"

// FlipCheck reverses the order of arguments for a check function.
func FlipCheck[T any](yield func(T) types.Checker[T]) func(T) types.Checker[T] {
	return func(base T) types.Checker[T] {
		return func(val T) bool {
			return yield(val)(base)
		}
	}
}

// FlipMap reverses the order of arguments for a map function.
func FlipMap[T any](yield func(T) types.Mapper[T]) func(T) types.Mapper[T] {
	return func(base T) types.Mapper[T] {
		return func(val T) T {
			return yield(val)(base)
		}
	}
}
