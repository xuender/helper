package list

// None returns a function that checks if no elements in a slice satisfy any provided predicate.
func None[T any](yields ...func(T) bool) func([]T) bool {
	return func(items []T) bool {
		for _, yield := range yields {
			for _, item := range items {
				if yield(item) {
					return false
				}
			}
		}

		return true
	}
}
