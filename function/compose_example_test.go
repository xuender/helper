package function_test

import (
	"fmt"

	"github.com/xuender/rg/function"
	"github.com/xuender/rg/math"
)

func ExampleCompose() {
	work := function.Compose(
		math.Add(2),
		math.Multiply(3),
	)

	fmt.Println(work(1))

	// Output:
	// 5
}
