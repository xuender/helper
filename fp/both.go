package fp

import "github.com/xuender/helper/gtype"

// Both returns a function that checks if all provided functions return true for the input value.
// The type parameter T can be of any type.
func Both[T any](checkers ...gtype.Checker[T]) gtype.Checker[T] {
	return func(val T) bool {
		for _, checker := range checkers {
			if !checker(val) {
				return false
			}
		}

		return true
	}
}
