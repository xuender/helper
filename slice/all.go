package slice

import "github.com/xuender/helper/types"

// All returns a function that checks if all elements in a slice satisfy all provided predicates.
func All[T any](chekders ...types.Checker[T]) func([]T) bool {
	return func(items []T) bool {
		for _, checker := range chekders {
			for _, item := range items {
				if !checker(item) {
					return false
				}
			}
		}

		return true
	}
}
