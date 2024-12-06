package slice_test

import (
	"fmt"

	"github.com/xuender/helper/slice"
)

func ExampleReduce() {
	adder := slice.Reduce(func(val1, val2 int) int {
		return val1 + val2
	})
	fmt.Println(adder([]int{1, 2, 3, 4}))

	// Output:
	// 10
}
