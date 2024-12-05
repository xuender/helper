package math

// Divide returns a function that divides the given dividend by the input divisor.
// The type parameter T must implement the Number interface.
func Divide[T Number](dividend T) func(T) T {
	return func(divisor T) T {
		return dividend / divisor
	}
}

// DivideReverse returns a function that divides the input dividend by the given divisor.
// The type parameter T must implement the Number interface.
func DivideReverse[T Number](divisor T) func(T) T {
	return func(dividend T) T {
		return dividend / divisor
	}
}
