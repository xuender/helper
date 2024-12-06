package fp

import "github.com/xuender/helper/types"

// Either returns true if any of the provided functions return true.
func Either[T any](checkers ...types.Checker[T]) types.Checker[T] {
	return func(i T) bool {
		for _, checker := range checkers {
			if checker(i) {
				return true
			}
		}

		return false
	}
}
