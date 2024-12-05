package function_test

import (
	"fmt"

	"github.com/xuender/rg/function"
)

func ExampleCurryReducer() {
	addOne := function.CurryReducer(func(val1, val2 int) int { return val1 + val2 })(1)

	fmt.Println(addOne(3))
	fmt.Println(addOne(7))

	// Output:
	// 4
	// 8
}
