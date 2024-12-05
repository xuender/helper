package list

// Any returns a function that checks if any element in a slice satisfies any provided predicate.
func Any[T any](yields ...func(T) bool) func([]T) bool {
	return func(items []T) bool {
		for _, yield := range yields {
			for _, item := range items {
				if yield(item) {
					return true
				}
			}
		}

		return false
	}
}
