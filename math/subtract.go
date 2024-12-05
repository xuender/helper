package math

// Subtract returns a function that subtracts the input value from the given base value.
// The type parameter T must implement the Number interface.
func Subtract[T Number](base T) func(T) T {
	return func(val T) T {
		return base - val
	}
}
