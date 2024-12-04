package logic

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
