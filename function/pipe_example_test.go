package function_test

import (
	"fmt"

	"github.com/xuender/ramdago/function"
	"github.com/xuender/ramdago/math"
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
