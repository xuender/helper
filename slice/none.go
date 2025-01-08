package slice

import "github.com/xuender/helper/types"

func None[S ~[]T, T any](items S, checkers ...types.Checker[T]) bool {
	for _, checker := range checkers {
		for _, item := range items {
			if checker(item) {
				return false
			}
		}
	}

	return true
}
