package slice

import "github.com/xuender/helper/types"

func Any[T any](items []T, checkers ...types.Checker[T]) bool {
	for _, checker := range checkers {
		for _, item := range items {
			if checker(item) {
				return true
			}
		}
	}

	return false
}
