package math_test

import (
	"fmt"

	"github.com/xuender/rg/math"
)

func ExampleDivide() {
	fmt.Println(math.Divide(71.0)(100))

	// Output:
	// 0.71
}

func ExampleDivideReverse() {
	fmt.Println(math.DivideReverse(100.0)(71))

	// Output:
	// 0.71
}
