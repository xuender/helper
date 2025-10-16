package cont

import "github.com/xuender/helper/gtype"

// None checks if no elements in the slice satisfy the provided checker function.
// It returns true if all elements fail the check, otherwise false.
func None[S ~[]T, T any](items S, checkers ...gtype.Checker[T]) bool {
	for _, checker := range checkers {
		for _, item := range items {
			if checker(item) {
				return false
			}
		}
	}

	return true
}
