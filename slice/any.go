package slice

import "github.com/xuender/helper/types"

// Any returns a function that checks if any element in a slice satisfies any provided predicate.
func Any[T any](checkers ...types.Checker[T]) func([]T) bool {
	return func(items []T) bool {
		for _, checker := range checkers {
			for _, item := range items {
				if checker(item) {
					return true
				}
			}
		}

		return false
	}
}
