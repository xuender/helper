package ops

// Result is a generic function that returns the provided argument unchanged.
// The variadic parameter is ignored and can be used for potential extensions.
func Result[R any](arg R, _ ...any) R {
	return arg
}

// Result1 is a generic function that acts as an identity function,
// returning the provided argument unchanged. The additional variadic
// parameters are ignored and can be used for future extensions.
func Result1[R any](arg R, _ ...any) R {
	return arg
}

// Result2 is a generic function that returns the second argument unchanged.
// It ignores the first argument and any additional variadic parameters.
// This function can be used as an identity function for the second argument provided.
func Result2[R any](_ any, arg R, _ ...any) R {
	return arg
}

// Result3 is a generic function that returns the third argument unchanged.
// It ignores the first two arguments and any additional variadic parameters.
// This function serves as an identity function for the third argument provided.
func Result3[R any](_, _ any, arg R, _ ...any) R {
	return arg
}

// Result4 is a generic function that returns the fourth argument unchanged.
// It ignores the first three arguments and any additional variadic parameters.
// This function serves as an identity function for the fourth argument provided.
func Result4[R any](_, _, _ any, arg R, _ ...any) R {
	return arg
}
