package logic

// Either returns true if any of the provided functions return true.
func Either[T any](yields ...func(T) bool) func(T) bool {
	return func(i T) bool {
		for _, yield := range yields {
			if yield(i) {
				return true
			}
		}

		return false
	}
}
