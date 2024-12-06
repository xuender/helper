package helper

import "github.com/xuender/helper/types"

// Either returns true if any of the provided functions return true.
func Either[T any](yields ...types.Checker[T]) types.Checker[T] {
	return func(i T) bool {
		for _, yield := range yields {
			if yield(i) {
				return true
			}
		}

		return false
	}
}
