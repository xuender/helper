package logic

// Both returns a function that checks if all provided functions return true for the input value.
// The type parameter T can be of any type.
func Both[T any](yields ...func(T) bool) func(T) bool {
	return func(val T) bool {
		for _, yield := range yields {
			if !yield(val) {
				return false
			}
		}

		return true
	}
}

// AllPass returns a function that checks if all provided functions return true for the input value.
// The type parameter T can be of any type.
//
// Deprecated: Use Both instead.
func AllPass[T any](yields ...func(T) bool) func(T) bool {
	return Both(yields...)
}
