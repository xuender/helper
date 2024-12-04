package math

func Multiply[T Number](base T) func(T) T {
	return func(val T) T {
		return base * val
	}
}
