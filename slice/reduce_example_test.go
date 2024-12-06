package slice_test

import (
	"fmt"

	"github.com/xuender/helper/slice"
)

func ExampleReduce() {
	sum := slice.Reduce([]int{1, 2, 3, 4}, func(val1, val2 int) int {
		return val1 + val2
	})

	fmt.Println(sum)

	// Output:
	// 10
}
