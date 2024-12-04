package function

func Pipe[T any](yields ...func(T) T) func(T) T {
	return func(val T) T {
		for _, yield := range yields {
			val = yield(val)
		}

		return val
	}
}
