package logic

func Both[T any](yields ...func(T) bool) func(T) bool {
	return func(i T) bool {
		for _, yield := range yields {
			if !yield(i) {
				return false
			}
		}

		return true
	}
}
