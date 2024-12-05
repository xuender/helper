package function_test

import (
	"fmt"

	"github.com/xuender/rg/function"
	"github.com/xuender/rg/math"
)

func ExamplePipe() {
	work := function.Pipe(
		math.Add(3),
		math.Multiply(2),
	)

	fmt.Println(work(4))

	// Output:
	// 14
}
