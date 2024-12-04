package list

func All[T any](yields ...func(T) bool) func([]T) bool {
	return func(items []T) bool {
		for _, yield := range yields {
			for _, item := range items {
				if !yield(item) {
					return false
				}
			}
		}

		return true
	}
}
